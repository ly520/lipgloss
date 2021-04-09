// +build windows

package lipgloss

import (
	"os"
	"sync"

	"golang.org/x/sys/windows"
)

var enableANSI sync.Once

// enableANSIColors enables support for ANSI color sequences in the Windows
// default console (cmd.exe and the PowerShell application). Note that this
// only works with Windows 10. Also note that Windows Terminal supports colors
// by default.
// enableANSIColors在Windows默认控制台中启用对ANSI颜色序列的支持(命令行以及PowerShell应用程序）。
// 请注意，这只适用于Windows10。另请注意，Windows终端默认支持颜色。
func enableLegacyWindowsANSI() {
	enableANSI.Do(func() {
		stdout := windows.Handle(os.Stdout.Fd())
		var originalMode uint32

		windows.GetConsoleMode(stdout, &originalMode)
		windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	})
}
