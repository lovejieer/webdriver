package webdriver

import (
	"github.com/tidwall/gjson"
	"io"
	"os"

	"os/exec"

	"errors"
)

type FirefoxDriver struct {
	*Session
	Cap Capabilities
	DriverPath    string
	BinaryPath string
	cmd     *exec.Cmd
	LogPath string
	LogFile string

	logFile *os.File

}

func NewFirefoxDriver(f *FirefoxDriver) *FirefoxDriver{

	return &FirefoxDriver{
		Session: f.Session,
		Cap: f.Capabilities,
		DriverPath: "",
		LogPath: "",
		BinaryPath: "",
		LogFile: "",
	}
}

func (f *FirefoxDriver) Start() error{

	csferr := "firefox driver start failed: "
	if f.cmd != nil {
		return errors.New(csferr + "firefoxdriver already running")
	}

	if f.LogPath != "" {
		//check if log-path is writable
		file, err := os.OpenFile(f.LogPath, os.O_WRONLY|os.O_CREATE, 0664)
		if err != nil {
			return errors.New(csferr + "unable to write in log path: " + err.Error())
		}
		file.Close()
	}

	var switches []string
	switches = append(switches, "-b "+f.BinaryPath)
	switches = append(switches, "-p 1234")

	f.cmd = exec.Command(f.DriverPath, switches...)
	stdout, err := f.cmd.StdoutPipe()
	if err != nil {
		return errors.New(csferr + err.Error())
	}
	stderr, err := f.cmd.StderrPipe()
	if err != nil {
		return errors.New(csferr + err.Error())
	}
	if err := f.cmd.Start(); err != nil {
		return errors.New(csferr + err.Error())
	}
	if f.LogFile != "" {
		flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
		f.logFile, err = os.OpenFile(f.LogFile, flags, 0640)
		if err != nil {
			return err
		}
		go io.Copy(f.logFile, stdout)
		go io.Copy(f.logFile, stderr)
	} else {
		go io.Copy(os.Stdout, stdout)
		go io.Copy(os.Stderr, stderr)
	}

	return nil
}

func (f *FirefoxDriver) Stop()error  {

		f.r.Get("http://127.0.0.1:1234/shutdown")
		err := f.cmd.Process.Signal(os.Interrupt)
		if f.logFile != nil {
			f.logFile.Close()
		}

		cmd := exec.Command("/usr/bin/pkill", "firefox")
		err = cmd.Run()
		if err != nil {
		return err
	}
		cmd2 := exec.Command("/usr/bin/killall", "firefox")
		err2 := cmd2.Run()
		if err2 != nil {
		return err

	}
		return nil

}

func (f *FirefoxDriver) NewSession() (*Session, error){

		if err :=f.Start(); err != nil{
			return nil, err
		}
		desired := f.Cap
		required := Capabilities{}
		p := params{"desiredCapabilities": desired, "requiredCapabilities": required}
       body := f.r.Post("http://127.0.0.1:1234/session",p)
		sid := gjson.Get(body, "value.sessionId")
		f.Id = sid.String()
       return f.Session,nil

}
