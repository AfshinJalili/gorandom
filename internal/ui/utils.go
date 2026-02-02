package ui

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func CopyToClipboard(text string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		// Prefer xclip; fallback to xsel if installed.
		cmd = exec.Command("xclip", "-selection", "clipboard")
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return fmt.Errorf("unsupported platform")
	}

	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}
	_, _ = in.Write([]byte(text))
	_ = in.Close()

	if err := cmd.Wait(); err != nil {
		// Fallback to xsel on linux if xclip is missing.
		if runtime.GOOS == "linux" && cmd.Path == "xclip" {
			return copyWithXsel(text)
		}
		return err
	}
	return nil
}

func copyWithXsel(text string) error {
	cmd := exec.Command("xsel", "--clipboard", "--input")
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	_, _ = in.Write([]byte(text))
	_ = in.Close()
	return cmd.Wait()
}
