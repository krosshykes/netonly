package netonly

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

var (
	advapi32                    = syscall.NewLazyDLL("advapi32.dll")
	procCreateProcessWithLogonW = advapi32.NewProc("CreateProcessWithLogonW")
)

const (
	// Use only network credentials for login
	LOGON_NETCREDENTIALS_ONLY uint32 = 0x00000002
	// The new process does not inherit the error mode of the calling process.
	// Instead, CreateProcessWithLogonW gives the new process the current
	// default error mode.
	CREATE_DEFAULT_ERROR_MODE uint32 = 0x04000000
	// Flag parameter that indicates to use the value set in ShowWindow
	STARTF_USESHOWWINDOW = 0x00000001
	// Tell OS wheather to show or not display the window
	ShowWindow = 1
)

// CreateProcessWithLogonW is a wrapper around the matching advapi32.dll
// function. This allows the running process to launch a process as a
// different user. It can also be used to stage credentials.
func CreateProcessWithLogonW(
	username *uint16,
	domain *uint16,
	password *uint16,
	logonFlags uint32,
	applicationName *uint16,
	commandLine *uint16,
	creationFlags uint32,
	environment *uint16,
	currentDirectory *uint16,
	startupInfo *syscall.StartupInfo,
	processInformation *syscall.ProcessInformation) error {
	r1, _, e1 := procCreateProcessWithLogonW.Call(
		uintptr(unsafe.Pointer(username)),
		uintptr(unsafe.Pointer(domain)),
		uintptr(unsafe.Pointer(password)),
		uintptr(logonFlags),
		uintptr(unsafe.Pointer(applicationName)),
		uintptr(unsafe.Pointer(commandLine)),
		uintptr(creationFlags),
		uintptr(unsafe.Pointer(environment)), // env
		uintptr(unsafe.Pointer(currentDirectory)),
		uintptr(unsafe.Pointer(startupInfo)),
		uintptr(unsafe.Pointer(processInformation)))
	runtime.KeepAlive(username)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(password)
	runtime.KeepAlive(applicationName)
	runtime.KeepAlive(commandLine)
	runtime.KeepAlive(environment)
	runtime.KeepAlive(currentDirectory)
	runtime.KeepAlive(startupInfo)
	runtime.KeepAlive(processInformation)
	if int(r1) == 0 {
		return os.NewSyscallError("CreateProcessWithLogonW", e1)
	}
	return nil
}

func StartProcess(user string, logonDomain string, pw string, path string, curDir string) error {

	username, _ := syscall.UTF16PtrFromString(user)
	domain, _ := syscall.UTF16PtrFromString(logonDomain)
	password, _ := syscall.UTF16PtrFromString(pw)
	logonFlags := LOGON_NETCREDENTIALS_ONLY
	applicationName, _ := syscall.UTF16PtrFromString(path)
	commandLine, _ := syscall.UTF16PtrFromString(``)
	creationFlags := CREATE_DEFAULT_ERROR_MODE
	currentDirectory, _ := syscall.UTF16PtrFromString(curDir)

	startupInfo := &syscall.StartupInfo{}
	startupInfo.ShowWindow = ShowWindow
	startupInfo.Flags = startupInfo.Flags | STARTF_USESHOWWINDOW
	processInfo := &syscall.ProcessInformation{}

	err := CreateProcessWithLogonW(
		username,
		domain,
		password,
		logonFlags,
		applicationName,
		commandLine,
		creationFlags,
		nil,
		currentDirectory,
		startupInfo,
		processInfo)
	return err
}
