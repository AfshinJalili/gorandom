package ui

import (
	"errors"
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

type fakeModel struct{}

func (fakeModel) Init() tea.Cmd                           { return nil }
func (fakeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return fakeModel{}, nil }
func (fakeModel) View() string                            { return "" }

type fakeTeaProgram struct {
	model     tea.Model
	err       error
	runCalled *bool
}

func (p *fakeTeaProgram) Run() (tea.Model, error) {
	if p.runCalled != nil {
		*p.runCalled = true
	}
	return p.model, p.err
}

func TestRunProgramUsesAltScreenAndRunsProgram(t *testing.T) {
	origFactory := newTeaProgram
	origAltScreen := withAltScreen
	t.Cleanup(func() {
		newTeaProgram = origFactory
		withAltScreen = origAltScreen
	})

	sentinel := tea.ProgramOption(func(*tea.Program) {})
	var gotOpts []tea.ProgramOption
	var runCalled bool

	withAltScreen = func() tea.ProgramOption {
		return sentinel
	}
	newTeaProgram = func(m tea.Model, opts ...tea.ProgramOption) teaProgram {
		gotOpts = opts
		return &fakeTeaProgram{
			model:     m,
			err:       nil,
			runCalled: &runCalled,
		}
	}

	in := fakeModel{}
	out, err := RunProgram(in)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if _, ok := out.(fakeModel); !ok {
		t.Fatalf("expected fakeModel output, got %T", out)
	}
	if len(gotOpts) != 1 {
		t.Fatalf("expected 1 program option, got %d", len(gotOpts))
	}
	if reflect.ValueOf(gotOpts[0]).Pointer() != reflect.ValueOf(sentinel).Pointer() {
		t.Fatalf("expected alt-screen option to be passed")
	}
	if !runCalled {
		t.Fatal("expected Run to be called")
	}
}

func TestRunProgramPropagatesRunError(t *testing.T) {
	origFactory := newTeaProgram
	t.Cleanup(func() { newTeaProgram = origFactory })

	wantErr := errors.New("run failed")
	newTeaProgram = func(m tea.Model, opts ...tea.ProgramOption) teaProgram {
		return &fakeTeaProgram{
			model: nil,
			err:   wantErr,
		}
	}

	_, err := RunProgram(fakeModel{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}
}
