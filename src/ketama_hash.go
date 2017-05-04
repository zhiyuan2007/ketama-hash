package main

import (
	"fmt"
	"hash/fnv"
	"sort"
	"strconv"
)

const MAXI = 1<<32 - 1

func hash_value(server string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(server))
	return h.Sum32()
}

func ketama_value(v, node_num uint32) []uint32 {
	step := MAXI / node_num
	node_v := make([]uint32, node_num)
	var i uint32
	for i = 0; i < node_num; i++ {
		node_v[i] = hash_value(strconv.FormatUint(uint64(v)+uint64(step)*uint64(i), 10))
	}
	return node_v
}

func sorted_map(server_hash_map map[uint32]string) []uint32 {
	all_sorted_hv := make([]uint32, 0)
	for i, _ := range server_hash_map {
		all_sorted_hv = append(all_sorted_hv, i)
	}
	sort.Slice(all_sorted_hv, func(i, j int) bool { return all_sorted_hv[i] < all_sorted_hv[j] })
	return all_sorted_hv
}

func binary_search_server(server_hash_map map[uint32]string, all_sorted_hv []uint32, key string) string {
	hash_v := hash_value(key)
	i := sort.Search(len(all_sorted_hv), func(i int) bool { return all_sorted_hv[i] >= hash_v })
	if i >= len(all_sorted_hv) {
		i = 0
	}
	//fmt.Println("key ", key, " i ", i, "hashv ", hash_v, " hash-nearest ", all_sorted_hv[i], " node ", server_hash_map[all_sorted_hv[i]])

	return server_hash_map[all_sorted_hv[i]]
}

func ketama_dispatch_result(servers []string, virtual_node uint32, testdata []string) map[string][]string {
	server_hash_map := make(map[uint32]string, 0)
	for _, v := range servers {
		x := hash_value(v)
		hash_arr := ketama_value(x, virtual_node)
		for _, iv := range hash_arr {
			server_hash_map[iv] = v
		}
	}

	all_sorted_hv := sorted_map(server_hash_map)
	//for i := range all_sorted_hv {
	//	fmt.Println(all_sorted_hv[i], server_hash_map[all_sorted_hv[i]])
	//}

	dispatch_result := make(map[string][]string, 0)
	for i := 0; i < len(testdata); i++ {
		node := binary_search_server(server_hash_map, all_sorted_hv, testdata[i])
		dispatch_result[node] = append(dispatch_result[node], testdata[i])
	}

	//{"1.1.1.1":[key1, key2, key9], "2.2.2.2":[key4, key8], "3.3.3.3":[key3, key5,key6, key7]}
	return dispatch_result
}

func main() {
	const TESTCOUNT = 1000
	var virtual_node uint32 = 200
	servers_begin := []string{"1.1.1.1", "2.2.2.2", "3.3.3.3"}
	servers_added := []string{"1.1.1.1", "2.2.2.2", "3.3.3.3", "4.4.4.4"}
	servers_reduce := []string{"1.1.1.1", "2.2.2.2"}
	testdata := make([]string, TESTCOUNT)
	for i := 0; i < TESTCOUNT; i++ {
		testdata[i] = strconv.Itoa(i) + "testdata" + strconv.Itoa(i)
	}
	result_begin := ketama_dispatch_result(servers_begin, virtual_node, testdata)
	for s, key_list := range result_begin {
		fmt.Println("begin server ", s, " has key count ", len(key_list))
		//for _, key := range key_list {
		//	fmt.Println("key: ", key)
		//}
	}
	result_added := ketama_dispatch_result(servers_added, virtual_node, testdata)
	for s, key_list := range result_added {
		fmt.Println("add server ", s, " has key count ", len(key_list))
		//for _, key := range key_list {
		//	fmt.Println("key: ", key)
		//}
	}
	result_reduce := ketama_dispatch_result(servers_reduce, virtual_node, testdata)
	for s, key_list := range result_reduce {
		fmt.Println("reduce server ", s, " has key count ", len(key_list))
		//for _, key := range key_list {
		//	fmt.Println("key: ", key)
		//}
	}

	fmt.Println("-----------add a new server------------hit ratio--------------------")
	for s, key_list := range result_begin {
		count := 0
		new_list := result_added[s]
		for _, key := range key_list {
			for _, newkey := range new_list {
				if key == newkey {
					count += 1
					break
				}
			}
		}
		var ratio float32 = float32(count*1.0) / float32(len(key_list))
		fmt.Printf("server %s, hit %d, total %d, hit_ratio %.2f\n", s, count, len(key_list), ratio)
	}

	fmt.Println("-----------reduce a server------------hit ratio--------------------")
	for s, key_list := range result_reduce {
		count := 0
		new_list := result_begin[s]
		for _, key := range key_list {
			for _, newkey := range new_list {
				if key == newkey {
					count += 1
					break
				}
			}
		}
		var ratio float32 = float32(count*1.0) / float32(len(key_list))
		fmt.Printf("server %s, hit %d, total %d, hit_ratio %.2f\n", s, count, len(key_list), ratio)
	}
}
