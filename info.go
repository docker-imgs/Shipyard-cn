/**
  shipyard.info :info数据结构
 */
package shipyard

type (
	ClusterInfo struct {
		Cpus           float64 `json:"cpus,omitempty"` //cpu定义
		Memory         float64 `json:"memory,omitempty"`//内存定义
		ContainerCount int     `json:"container_count,omitempty"`//容器总数
		EngineCount    int     `json:"engine_count,omitempty"`//引擎总数
		ImageCount     int     `json:"image_count,omitempty"` //镜像总数
		ReservedCpus   float64 `json:"reserved_cpus,omitempty"`//分配的cpus
		ReservedMemory float64 `json:"reserved_memory,omitempty"`//分配的内存
		Version        string  `json:"version,omitempty"` //版本
	}
)
