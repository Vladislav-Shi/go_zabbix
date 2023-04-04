package app

import (
	"fmt"
	"log"
	"reflect"

	cpu "github.com/shirou/gopsutil/v3/cpu"
	host "github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	ps "github.com/shirou/gopsutil/v3/process"
)

type ProcessInfo struct {
	cpuPercent    float64 // CPU_Percent returns how many percent of the CPU time this process uses
	memoryPercent float32
	background    bool
	pid           int32
	createTime    int64
	cwd           string // Cwd returns current working directory of the process
	exe           string // Exe returns executable path of the process.
	name          string
	ppid          int32 // parent pid
	username      string
	status        string // Return value could be one of these. R: Running S: Sleep T: Stop I: Idle Z: Zombie W: Wait L: Lock The character is same within all supported platforms.
}

func (p *ProcessInfo) ToMap(executes []string) map[string]string {
	/**
	формирует из структуры карту.
	Параметр нужен чтобы извлекать ненужные поля из итоговой карты. Передаются названия полей
	*/
	m := make(map[string]string)
	var flag bool
	for i := 0; i < 11; i++ {
		fieldName := reflect.TypeOf(*p).Field(i).Name
		flag = true
		for _, b := range executes {
			if fieldName == b {
				flag = false
			}
		}
		if flag {
			a := reflect.ValueOf(*p).Field(i)
			m[fieldName] = fmt.Sprintf("%v", a)
		}
	}
	return m
}

func GetProcessesInfo() []ProcessInfo {
	/*
		Возвращает список процессов в виде строк
	*/
	processes, err := ps.Processes()
	var processesInfo []ProcessInfo
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
	}
	for _, process := range processes {
		name, _ := process.Name()
		username, _ := process.Username()
		cwd, _ := process.Cwd()
		parent, _ := process.Parent()
		cpuPercent, _ := process.CPUPercent()
		memoryPercent, _ := process.MemoryPercent()
		background, _ := process.Background()
		createTime, _ := process.CreateTime()
		exe, _ := process.Exe()
		statuses, _ := process.Status()
		processesInfo = append(processesInfo, ProcessInfo{
			cpuPercent:    cpuPercent,
			memoryPercent: memoryPercent,
			background:    background,
			pid:           process.Pid,
			createTime:    createTime,
			cwd:           cwd,
			exe:           exe,
			name:          name,
			ppid:          parent.Pid,
			username:      username,
			status:        statuses[0],
		})
	}
	return processesInfo
}

func GetProcessInfo(pid int32) ProcessInfo {
	process, err := ps.NewProcess(pid)
	if err != nil {
		log.Println(err)
	}
	name, _ := process.Name()
	username, _ := process.Username()
	cwd, _ := process.Cwd()
	parent, _ := process.Parent()
	cpuPercent, _ := process.CPUPercent()
	memoryPercent, _ := process.MemoryPercent()
	background, _ := process.Background()
	createTime, _ := process.CreateTime()
	exe, _ := process.Exe()
	statuses, _ := process.Status()

	processInfo := ProcessInfo{
		cpuPercent:    cpuPercent,
		memoryPercent: memoryPercent,
		background:    background,
		pid:           process.Pid,
		createTime:    createTime,
		cwd:           cwd,
		exe:           exe,
		name:          name,
		ppid:          parent.Pid,
		username:      username,
		status:        statuses[0],
	}
	return processInfo
}

func GetMemoryInfo() map[string]string {
	v, _ := mem.VirtualMemory()
	m := make(map[string]string)
	m["Total"] = fmt.Sprintf("%v", v.Total)
	m["Free"] = fmt.Sprintf("%v", v.Free)
	m["UsedPercent"] = fmt.Sprintf("%v", v.UsedPercent)
	m["Cached"] = fmt.Sprintf("%v", v.Cached)
	m["SwapTotal"] = fmt.Sprintf("%v", v.SwapTotal)
	m["SwapFree"] = fmt.Sprintf("%v", v.SwapFree)
	return m
}

func GetCpuStat() []map[string]string {
	/*
		Информация о процессе
		CPUInfo on linux will return 1 item per physical thread.
	*/
	var m []map[string]string
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Println(err)
	}
	for i, v := range cpuInfo {
		m = append(m, make(map[string]string))
		m[i]["CPU"] = fmt.Sprintf("%v", v.CPU)
		m[i]["CacheSize"] = fmt.Sprintf("%v", v.CacheSize)
		m[i]["Model"] = fmt.Sprintf("%v", v.Model)
		m[i]["Cores"] = fmt.Sprintf("%v", v.Cores)
		m[i]["ModelName"] = fmt.Sprintf("%v", v.ModelName)
		m[i]["Mhz"] = fmt.Sprintf("%v", v.Mhz)
	}
	return m
}

func GetHostInfo() map[string]string {
	m := make(map[string]string)
	hostInfo, err := host.Info()
	if err != nil {
		log.Println(err)
	}
	m["HostID"] = fmt.Sprintf("%v", hostInfo.HostID)
	m["Hostname"] = fmt.Sprintf("%v", hostInfo.Hostname)
	m["OS"] = fmt.Sprintf("%v", hostInfo.OS)
	m["Platform"] = fmt.Sprintf("%v", hostInfo.Platform)
	m["Procs"] = fmt.Sprintf("%v", hostInfo.Procs)
	m["KernelVersion"] = fmt.Sprintf("%v", hostInfo.KernelVersion)
	m["BootTime"] = fmt.Sprintf("%v", hostInfo.BootTime)

	return m
}

func ProcessStop(pid int32) bool {
	/**
	*
	 */
	process, err := ps.NewProcess(pid)
	if err != nil {
		return false
	}
	err = process.Suspend()
	if err != nil {
		return false
	}
	return true
}

func ProcessKill(pid int32) bool {
	/**
	*
	 */
	process, err := ps.NewProcess(pid)
	if err != nil {
		return false
	}
	err = process.Kill()
	if err != nil {
		return false
	}
	return true
}

func ProcessResume(pid int32) bool {
	/**
	* Отправляет SIGCONT процессу для возобновления работы
	 */
	process, err := ps.NewProcess(pid)
	if err != nil {
		return false
	}
	err = process.Resume()
	if err != nil {
		return false
	}
	return true
}

// TODO: Для баш скриптов пока отложить, мб не надо
func CopliteBashScript(bashFileName string) {
	/**
	*
	 */
}

func GetBashScript(bashFileName string) {
	/**
	*
	 */
}

func CreateBashScript(bashFileName string) {
	/**
	*
	 */
}
