package service

import (
	"github.com/shirou/gopsutil/host"
	"gogo/backend/app/dto"
)

type DashboardService struct{}

type IDashboardService interface {
	LoadOsInfo() (*dto.OsInfo, error)
}

func NewIDashboardService() IDashboardService {
	return &DashboardService{}
}

func (u *DashboardService) LoadOsInfo() (*dto.OsInfo, error) {
	var baseInfo dto.OsInfo
	hostInfo, err := host.Info()
	if err != nil {
		return nil, err
	}
	baseInfo.OS = hostInfo.OS
	baseInfo.Platform = hostInfo.Platform
	baseInfo.PlatformFamily = hostInfo.PlatformFamily
	baseInfo.KernelArch = hostInfo.KernelArch
	baseInfo.KernelVersion = hostInfo.KernelVersion

	//diskInfo, err := disk.Usage(global.CONF.System.BaseDir)
	//if err == nil {
	//	baseInfo.DiskSize = int64(diskInfo.Free)
	//}

	if baseInfo.KernelArch == "armv7l" {
		baseInfo.KernelArch = "armv7"
	}
	if baseInfo.KernelArch == "x86_64" {
		baseInfo.KernelArch = "amd64"
	}

	return &baseInfo, nil
}
