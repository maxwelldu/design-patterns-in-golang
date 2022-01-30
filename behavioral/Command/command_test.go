package Command
import "testing"

// client
func TestCommand_Execute(t *testing.T) {
	tv := &Tv{}

	onCommand := &onCommand{device: tv}
	offCommand := &offCommand{device: tv}

	onButton := &button{command: onCommand}
	onButton.press()

	offButton := &button{command: offCommand}
	offButton.press()
}
