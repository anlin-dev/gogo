package dto

type OsInfo struct {
	OS             string `json:"os"`
	Platform       string `json:"platform"`
	PlatformFamily string `json:"platformFamily"`
	KernelArch     string `json:"kernelArch"`
	KernelVersion  string `json:"kernelVersion"`

	DiskSize int64 `json:"diskSize"`
}
