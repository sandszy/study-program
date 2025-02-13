package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

// chapter1:定义数据结构
type Host struct {
	ID           int
	RemainingCPU int
	RemainingMem int
}

type Container struct {
	CPU int
	Mem int
}

type Scheduler interface {
	Schedule(hosts []*Host, c *Container) *Host
	Name() string
}

// chapter2:实现调度算法

// First Fit 算法
type FirstFit struct{}

func (s *FirstFit) Schedule(hosts []*Host, c *Container) *Host {
	for _, h := range hosts {
		if h.RemainingCPU >= c.CPU && h.RemainingMem >= c.Mem {
			return h
		}
	}
	return nil
}
func (s *FirstFit) Name() string { return "First Fit" }

// Best Fit 算法
type BestFit struct{}

func (s *BestFit) Schedule(hosts []*Host, c *Container) *Host {
	var bestHost *Host
	minSum := math.MaxInt
	for _, h := range hosts {
		if h.RemainingCPU >= c.CPU && h.RemainingMem >= c.Mem {
			sum := (h.RemainingCPU - c.CPU) + (h.RemainingMem - c.Mem)
			if sum < minSum {
				minSum = sum
				bestHost = h
			}
		}
	}
	return bestHost
}
func (s *BestFit) Name() string { return "Best Fit" }

// Worst Fit 算法
type WorstFit struct{}

func (s *WorstFit) Schedule(hosts []*Host, c *Container) *Host {
	var worstHost *Host
	maxSum := -1
	for _, h := range hosts {
		if h.RemainingCPU >= c.CPU && h.RemainingMem >= c.Mem {
			sum := (h.RemainingCPU - c.CPU) + (h.RemainingMem - c.Mem)
			if sum > maxSum {
				maxSum = sum
				worstHost = h
			}
		}
	}
	return worstHost
}
func (s *WorstFit) Name() string { return "Worst Fit" }

// Random 算法
type Random struct {
	rand *rand.Rand
}

func NewRandom() *Random {
	return &Random{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
func (s *Random) Schedule(hosts []*Host, c *Container) *Host {
	var candidates []*Host
	for _, h := range hosts {
		if h.RemainingCPU >= c.CPU && h.RemainingMem >= c.Mem {
			candidates = append(candidates, h)
		}
	}
	if len(candidates) == 0 {
		return nil
	}
	return candidates[s.rand.Intn(len(candidates))]
}
func (s *Random) Name() string { return "Random" }

// chapter3:模拟器与结果统计
type Simulator struct {
	Hosts      []*Host
	Containers []*Container
}

type Result struct {
	Algorithm      string
	Success        int
	CPUUtilization float64
	MemUtilization float64
	HostsUsed      int
}

func (sim *Simulator) Run(scheduler Scheduler) Result {
	// 重置资源
	for _, h := range sim.Hosts {
		h.RemainingCPU = 128
		h.RemainingMem = 1024
	}

	success := 0
	for _, c := range sim.Containers {
		h := scheduler.Schedule(sim.Hosts, c)
		if h != nil {
			h.RemainingCPU -= c.CPU
			h.RemainingMem -= c.Mem
			success++
		}
	}

	// 计算指标
	totalCPU := len(sim.Hosts) * 128
	totalMem := len(sim.Hosts) * 1024
	usedCPU, usedMem := 0, 0
	hostsUsed := 0
	for _, h := range sim.Hosts {
		usedCPU += 128 - h.RemainingCPU
		usedMem += 1024 - h.RemainingMem
		if h.RemainingCPU < 128 || h.RemainingMem < 1024 {
			hostsUsed++
		}
	}

	return Result{
		Algorithm:      scheduler.Name(),
		Success:        success,
		CPUUtilization: float64(usedCPU) / float64(totalCPU),
		MemUtilization: float64(usedMem) / float64(totalMem),
		HostsUsed:      hostsUsed,
	}
}

// chapter4：主函数与测试数据

func main() {
	rand.Seed(time.Now().UnixNano())

	// 初始化500台宿主机
	hosts := make([]*Host, 500)
	for i := 0; i < 500; i++ {
		hosts[i] = &Host{ID: i + 1, RemainingCPU: 128, RemainingMem: 1024}
	}

	// 生成20000个容器请求（正态分布）
	containers := make([]*Container, 20000)
	for i := 0; i < 20000; i++ {
		cpu := int(math.Round(rand.NormFloat64()*8 + 16))
		if cpu < 1 {
			cpu = 1
		} else if cpu > 32 {
			cpu = 32
		}

		mem := int(math.Round(rand.NormFloat64()*32 + 64))
		if mem < 4 {
			mem = 4
		} else if mem > 128 {
			mem = 128
		}
		containers[i] = &Container{CPU: cpu, Mem: mem}
	}

	sim := &Simulator{Hosts: hosts, Containers: containers}

	// 调度算法列表
	schedulers := []Scheduler{
		&FirstFit{},
		&BestFit{},
		&WorstFit{},
		NewRandom(),
	}

	// 运行模拟
	var results []Result
	for _, scheduler := range schedulers {
		result := sim.Run(scheduler)
		results = append(results, result)
	}

	// 打印结果表格
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Algorithm", "Success", "CPU Util (%)", "Mem Util (%)", "Hosts Used"})

	for _, res := range results {
		row := []string{
			res.Algorithm,
			fmt.Sprintf("%d", res.Success),
			fmt.Sprintf("%.2f%%", res.CPUUtilization*100),
			fmt.Sprintf("%.2f%%", res.MemUtilization*100),
			fmt.Sprintf("%d/500", res.HostsUsed),
		}
		table.Append(row)
	}
	table.Render()
}
