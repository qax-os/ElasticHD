package search

// Params 前端传递过来的参数
type Params struct {
	Serverhost string `json:"host"`
	Indices    string `json:"indices"`
	Types      string `json:"types"`
	Method     string `json:"method"`
	Options    string `json:"options"`
	Body       string `json:"body"`
	SQL        string `json:"sql"`
	Index      string `json:"index"`
}

// ReturnResult 返回的结果
type ReturnResult struct {
	Result int         `json:"result"` // 0：success 1:faild
	Data   interface{} `json:"data"`   // 返回的data
}

// Indices ...
type Indices struct {
	IndexOrder string
	Health     string `json:"health"`
	Status     string `json:"status"`
	Index      string `json:"index"`
	UUID       string `json:"uuid"`
	Pri        string `json:"pri"`
	Rep        string `json:"rep"`
	Count      string `json:"docs.count"`
	Delete     string `json:"docs.deleted"`
	Size       string `json:"store.size"`
	PriSize    string `json:"pri.store.size"`
}

// Stats ...
type Stats struct {
	ClusterName      string      `json:"cluster_name"` // 集群名称
	Status           string      `json:"status"`       // 集群状态
	Shards           int64       `json:"shards"`       // 分片数
	Docs             int64       `json:"docs"`         // 文档数
	Size             string      `json:"size"`         // 内存使用
	Indices          int64       `json:"indices"`      // 索引数
	Versions         []string    `json:"versions"`     // elasticsearch 版本们
	Systems          string      `json:"systems"`      // 系统类型
	MemToatl         int64       `json:"memtotal"`     // 内存大小
	MemFree          int64       `json:"memfree"`      // 内存空闲
	MemUsed          int64       `json:"memused"`      // 内存占用
	FsTotal          int64       `json:"fstotal"`      // 磁盘大小
	FsFree           int64       `json:"fsfree"`       // 磁盘空闲
	FsUsed           int64       `json:"fsused"`       // 磁盘占用
	Plugins          interface{} `json:"plugins"`      // 插件列表
	JvmThreads       int64       `json:"jvm_threads"`  // jvm threads
	JvmMaxUptime     string      `json:"jvm_max_uptime"`
	JvmMemUsed       int64       `json:"jvm_mem_used"`
	JvmMemMaxUse     int64       `json:"jvm_mem_max_use"`
	JvmVersions      string      `json:"jvm_versions"`
	FieldDataMemUsed int64       `json:"fielddata_mem_used"`
	QueryCache       interface{} `json:"query_cache"`
	Timestamp        int64       `json:"timestamp"`
	CPUUsed          int         `json:"cpuused"`
	CPUFree          int         `json:"cpufree"`
}

// System ...
type System struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// JVMVersion ...
type JVMVersion struct {
	Version   string `json:"version"`
	VMName    string `json:"vm_name"`
	VMVersion string `json:"vm_version"`
	VMVendor  string `json:"vm_vendor"`
	Count     int    `json:"count"`
}

// ClusterStats 集群状态
type ClusterStats struct {
	Timestamp   int64  `json:"timestamp"`
	ClusterName string `json:"cluster_name"`
	Status      string `json:"status"`
	Indices     struct {
		Count  int64 `json:"count"`
		Shards struct {
			Total       int64   `json:"total"`
			Primaries   int64   `json:"primaries"`
			Replication float64 `json:"replication"`
			// Index       struct {
			// 	Shards struct {
			// 		Min int `json:"min"`
			// 		Max int `json:"max"`
			// 		Avg int `json:"avg"`
			// 	} `json:"shards"`
			// 	Primaries struct {
			// 		Min int `json:"min"`
			// 		Max int `json:"max"`
			// 		Avg int `json:"avg"`
			// 	} `json:"primaries"`
			// 	Replication struct {
			// 		Min int `json:"min"`
			// 		Max int `json:"max"`
			// 		Avg int `json:"avg"`
			// 	} `json:"replication"`
			// } `json:"index"`
		} `json:"shards"`
		Docs struct {
			Count   int64 `json:"count"`
			Deleted int64 `json:"deleted"`
		} `json:"docs"`
		Store struct {
			Size                 string `json:"size"`
			SizeInBytes          int    `json:"size_in_bytes"`
			ThrottleTime         string `json:"throttle_time"`
			ThrottleTimeInMillis int    `json:"throttle_time_in_millis"`
		} `json:"store"`
		Fielddata struct {
			MemorySize        string `json:"memory_size"`
			MemorySizeInBytes int64  `json:"memory_size_in_bytes"`
			Evictions         int    `json:"evictions"`
		} `json:"fielddata"`
		QueryCache struct {
			MemorySize        string `json:"memory_size"`
			MemorySizeInBytes int    `json:"memory_size_in_bytes"`
			TotalCount        int    `json:"total_count"`
			HitCount          int    `json:"hit_count"`
			MissCount         int    `json:"miss_count"`
			CacheSize         int    `json:"cache_size"`
			CacheCount        int    `json:"cache_count"`
			Evictions         int    `json:"evictions"`
		} `json:"query_cache"`
		Completion struct {
			Size        string `json:"size"`
			SizeInBytes int    `json:"size_in_bytes"`
		} `json:"completion"`
		Segments struct {
			Count                     int    `json:"count"`
			Memory                    string `json:"memory"`
			MemoryInBytes             int    `json:"memory_in_bytes"`
			TermsMemory               string `json:"terms_memory"`
			TermsMemoryInBytes        int    `json:"terms_memory_in_bytes"`
			StoredFieldsMemory        string `json:"stored_fields_memory"`
			StoredFieldsMemoryInBytes int    `json:"stored_fields_memory_in_bytes"`
			TermVectorsMemory         string `json:"term_vectors_memory"`
			TermVectorsMemoryInBytes  int    `json:"term_vectors_memory_in_bytes"`
			NormsMemory               string `json:"norms_memory"`
			NormsMemoryInBytes        int    `json:"norms_memory_in_bytes"`
			DocValuesMemory           string `json:"doc_values_memory"`
			DocValuesMemoryInBytes    int    `json:"doc_values_memory_in_bytes"`
			IndexWriterMemory         string `json:"index_writer_memory"`
			IndexWriterMemoryInBytes  int    `json:"index_writer_memory_in_bytes"`
			VersionMapMemory          string `json:"version_map_memory"`
			VersionMapMemoryInBytes   int    `json:"version_map_memory_in_bytes"`
			FixedBitSet               string `json:"fixed_bit_set"`
			FixedBitSetMemoryInBytes  int    `json:"fixed_bit_set_memory_in_bytes"`
			FileSizes                 struct {
			} `json:"file_sizes"`
		} `json:"segments"`
		Percolator struct {
			NumQueries int `json:"num_queries"`
		} `json:"percolator"`
	} `json:"indices"`
	Nodes struct {
		Count struct {
			Total            int `json:"total"`
			Data             int `json:"data"`
			CoordinatingOnly int `json:"coordinating_only"`
			Master           int `json:"master"`
			Ingest           int `json:"ingest"`
		} `json:"count"`
		Versions []string `json:"versions"`
		Os       struct {
			AvailableProcessors int `json:"available_processors"`
			AllocatedProcessors int `json:"allocated_processors"`
			Names               []struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
			} `json:"names"`
			Mem struct {
				Total        string `json:"total"`
				TotalInBytes int64  `json:"total_in_bytes"`
				Free         string `json:"free"`
				FreeInBytes  int64  `json:"free_in_bytes"`
				Used         string `json:"used"`
				UsedInBytes  int64  `json:"used_in_bytes"`
				FreePercent  int    `json:"free_percent"`
				UsedPercent  int    `json:"used_percent"`
			} `json:"mem"`
		} `json:"os"`
		Process struct {
			CPU struct {
				Percent int `json:"percent"`
			} `json:"cpu"`
			OpenFileDescriptors struct {
				Min int `json:"min"`
				Max int `json:"max"`
				Avg int `json:"avg"`
			} `json:"open_file_descriptors"`
		} `json:"process"`
		Jvm struct {
			MaxUptime         string `json:"max_uptime"`
			MaxUptimeInMillis int    `json:"max_uptime_in_millis"`
			Versions          []struct {
				Version   string `json:"version"`
				VMName    string `json:"vm_name"`
				VMVersion string `json:"vm_version"`
				VMVendor  string `json:"vm_vendor"`
				Count     int64  `json:"count"`
			} `json:"versions"`
			Mem struct {
				HeapUsed        string `json:"heap_used"`
				HeapUsedInBytes int64  `json:"heap_used_in_bytes"`
				HeapMax         string `json:"heap_max"`
				HeapMaxInBytes  int64  `json:"heap_max_in_bytes"`
			} `json:"mem"`
			Threads int64 `json:"threads"`
		} `json:"jvm"`
		Fs struct {
			Total            string `json:"total"`
			TotalInBytes     int64  `json:"total_in_bytes"`
			Free             string `json:"free"`
			FreeInBytes      int64  `json:"free_in_bytes"`
			Available        string `json:"available"`
			AvailableInBytes int64  `json:"available_in_bytes"`
		} `json:"fs"`
		Plugins []struct {
			Name        string `json:"name"`
			Version     string `json:"version"`
			Description string `json:"description"`
			Classname   string `json:"classname"`
		} `json:"plugins"`
	} `json:"nodes"`
}
