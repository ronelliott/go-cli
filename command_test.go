package cli

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCommandString_Empty(t *testing.T) {
	name, idx := getCommandName([]string{})
	assert.Equal(t, "", name)
	assert.Equal(t, -1, idx)
}

func TestGetCommandString_Single(t *testing.T) {
	name, idx := getCommandName([]string{"foo"})
	assert.Equal(t, "foo", name)
	assert.Equal(t, 0, idx)
}

func TestGetCommandString_Multiple(t *testing.T) {
	name, idx := getCommandName([]string{"foo", "bar"})
	assert.Equal(t, "foo", name)
	assert.Equal(t, 0, idx)
}

func TestGetCommandString_WithFlags(t *testing.T) {
	name, idx := getCommandName([]string{"-f", "--bar", "foo"})
	assert.Equal(t, "foo", name)
	assert.Equal(t, 2, idx)
}

func TestNew_SetName(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	assert.Equal(t, "go-cli.test", cmd.Name)
}

func TestNew_SetDescription(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	assert.Equal(t, "bar", cmd.Description)
}

func TestNew_NoSubs(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(cmd.Subs))
}

type TestNew_ErrorCmd struct {
	t       *testing.T
	Verbose bool `default:"Nein" short:"v"`
}

func (opts *TestNew_ErrorCmd) Run() error {
	return nil
}

func TestNew_Error(t *testing.T) {
	cmd, err := New("bar", &TestNew_ErrorCmd{t: t})
	assert.NotNil(t, err)
	assert.Nil(t, cmd)
}

type TestNewSub_ErrorCmd struct {
	t       *testing.T
	Verbose bool `default:"Nein" short:"v"`
}

func (opts *TestNewSub_ErrorCmd) Run() error {
	return nil
}

func TestNewSub_Error(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.NotNil(t, cmd)
	assert.Nil(t, err)
	sub, err := cmd.NewSub("foo", "bar", "dar", &TestNewSub_ErrorCmd{t: t})
	assert.NotNil(t, err)
	assert.Nil(t, sub)
}

func TestNewSub_ParentSubs(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	cmd.NewSub("car", "dar", "dar", nil)
	assert.Equal(t, 1, len(cmd.Subs))
}

func TestNewSub_SetName(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	sub, err := cmd.NewSub("car", "dar", "dar", nil)
	assert.Nil(t, err)
	assert.Equal(t, "car", sub.Name)
}

func TestNewSub_SetDescription(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	sub, err := cmd.NewSub("car", "dar", "dar", nil)
	assert.Nil(t, err)
	assert.Equal(t, "dar", sub.Description)
}

func TestNewSub_NoSubs(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	sub, err := cmd.NewSub("foo", "bar", "bar", nil)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(sub.Subs))
}

func TestHasCallback_NoCallback(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	assert.False(t, cmd.HasCallback())
}

type TestHasCallback_WithCallbackCmd struct{}

func (opts *TestHasCallback_WithCallbackCmd) Run() error {
	return nil
}

func TestHasCallback_WithCallback(t *testing.T) {
	cmd, err := New("bar", &TestHasCallback_WithCallbackCmd{})
	assert.Nil(t, err)
	assert.True(t, cmd.HasCallback())
}

func TestHasSubs_NoSubs(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	assert.False(t, cmd.HasSubs())
}

func TestHasSubs_WithSubs(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	cmd.NewSub("foo", "bar", "bar", nil)
	assert.True(t, cmd.HasSubs())
}

func TestHasSub_NoSubs(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	assert.False(t, cmd.HasSub("bar"))
}

func TestHasSub_NoChild(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	cmd.NewSub("foo", "bar", "bar", nil)
	assert.False(t, cmd.HasSub("bar"))
}

func TestHasSub_WithChild(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	cmd.NewSub("foo", "bar", "bar", nil)
	assert.True(t, cmd.HasSub("foo"))
}

var TestRun_ArgsCalled bool

type TestRun_ArgsCmd struct {
	t       *testing.T
	Verbose bool `short:"v" description:"Show verbose debug information"`
}

func (opts *TestRun_ArgsCmd) Run() error {
	TestRun_ArgsCalled = true
	assert.True(opts.t, opts.Verbose)
	return nil
}

func TestRun_WithArgs(t *testing.T) {
	assert.False(t, TestRun_ArgsCalled)
	cmd, err := New("bar", &TestRun_ArgsCmd{t: t})
	assert.Nil(t, err)
	err = cmd.Run([]string{"-v"})
	assert.Nil(t, err)
	assert.True(t, TestRun_ArgsCalled)
}

var TestRun_LeftoverArgsCalled bool

type TestRun_LeftoverArgsCmd struct {
	t       *testing.T
	Args    []string `positional:"true"`
	Verbose bool     `short:"v" description:"Show verbose debug information"`
}

func (opts *TestRun_LeftoverArgsCmd) Run() error {
	TestRun_LeftoverArgsCalled = true
	assert.Equal(opts.t, []string{"foo", "bar"}, opts.Args)
	return nil
}

func TestRun_LeftoverArgs(t *testing.T) {
	assert.False(t, TestRun_LeftoverArgsCalled)
	cmd, err := New("bar", &TestRun_LeftoverArgsCmd{t: t})
	assert.Nil(t, err)
	err = cmd.Run([]string{"-v", "foo", "bar"})
	assert.Nil(t, err)
	assert.True(t, TestRun_LeftoverArgsCalled)
}

func TestRun_NoCallback(t *testing.T) {
	cmd, err := New("bar", nil)
	assert.Nil(t, err)
	err = cmd.Run(nil)
	assert.NotNil(t, err)
}

var TestRun_ErrorCalled bool

type TestRun_ErrorCmd struct {
	t       *testing.T
	Verbose bool `short:"v" description:"Show verbose debug information"`
}

func (opts *TestRun_ErrorCmd) Run() error {
	TestRun_ErrorCalled = true
	assert.True(opts.t, opts.Verbose)
	return nil
}

func TestRun_Error(t *testing.T) {
	assert.False(t, TestRun_ErrorCalled)
	cmd, err := New("bar", &TestRun_ErrorCmd{t: t})
	assert.Nil(t, err)
	err = cmd.Run([]string{"-z"})
	assert.NotNil(t, err)
	assert.False(t, TestRun_ErrorCalled)
}

var TestSubRun_LeftoverArgsCalled bool

type TestSubRun_LeftoverArgsCmd struct {
	t    *testing.T
	Args []string `positional:"true"`
}

func (opts *TestSubRun_LeftoverArgsCmd) Run() error {
	TestSubRun_LeftoverArgsCalled = true
	assert.Equal(opts.t, []string{"car", "dar", "far"}, opts.Args)
	return nil
}

func TestSubRun_LeftoverArgs(t *testing.T) {
	assert.False(t, TestSubRun_LeftoverArgsCalled)
	cmd, err := New("foo", nil)
	assert.Nil(t, err)
	cmd.NewSub("bar", "dar", "dar", &TestSubRun_LeftoverArgsCmd{t: t})
	err = cmd.Run([]string{"bar", "car", "dar", "far"})
	assert.Nil(t, err)
	assert.True(t, TestSubRun_LeftoverArgsCalled)
}

var TestSubRun_CallCallbackCalled bool

type TestSubRun_CallCallbackCmd struct {
	t *testing.T
}

func (opts *TestSubRun_CallCallbackCmd) Run() error {
	TestSubRun_CallCallbackCalled = true
	return nil
}

func TestSubRun_CallCallback(t *testing.T) {
	assert.False(t, TestSubRun_CallCallbackCalled)
	cmd, err := New("foo", nil)
	assert.Nil(t, err)
	cmd.NewSub("bar", "dar", "dar", &TestSubRun_CallCallbackCmd{t: t})
	err = cmd.Run([]string{"bar"})
	assert.Nil(t, err)
	assert.True(t, TestSubRun_CallCallbackCalled)
}

func TestSubRun_NoCallback(t *testing.T) {
	cmd, err := New("foo", nil)
	assert.Nil(t, err)
	cmd.NewSub("bar", "dar", "dar", nil)
	err = cmd.Run([]string{"bar"})
	assert.NotNil(t, err)
}

var TestNestedSubRun_CallCallbackCalled bool

type TestNestedSubRun_CallCallbackCmd struct {
	t *testing.T
}

func (opts *TestNestedSubRun_CallCallbackCmd) Run() error {
	TestNestedSubRun_CallCallbackCalled = true
	return nil
}

func TestNestedSubRun_CallCallback(t *testing.T) {
	assert.False(t, TestNestedSubRun_CallCallbackCalled)
	cmd, err := New("foo", nil)
	assert.Nil(t, err)
	sub, err := cmd.NewSub("bar", "dar", "dar", nil)
	assert.Nil(t, err)
	sub.NewSub("far", "car", "car", &TestNestedSubRun_CallCallbackCmd{t: t})
	err = cmd.Run([]string{"bar", "far"})
	assert.Nil(t, err)
	assert.True(t, TestNestedSubRun_CallCallbackCalled)
}

type TestWriteHelpCmd struct {
	Name string `
        default:"foo"
        description:"The name to use"
        help:"What do you want to name this thing?"
        long:"name"
        short:"n"`

	Verbose bool `
        default:"false"
        description:"Use verbose logging."
        help:"Be very talkative when logging"
        long:"verbose"
        short:"v"`
}

func (opts *TestWriteHelpCmd) Run() error {
	return nil
}

func TestWriteHelp(t *testing.T) {
	cmd, err := New("foo", &TestWriteHelpCmd{})
	assert.Nil(t, err)
	buf := bytes.Buffer{}
	cmd.WriteHelp(&buf)

	expected := "  -n string\n    \tThe name to use (default \"foo\")\n  " +
		"-name string\n    \tThe name to use (default \"foo\")\n  " +
		"-v\tUse verbose logging.\n  -verbose\n    \tUse verbose " +
		"logging.\n"

	assert.Equal(t, expected, buf.String())
}
