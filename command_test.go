package cli

import (
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
    cmd := New("bar", nil)
    assert.Equal(t, "go-cli.test", cmd.Name)
}

func TestNew_SetDescription(t *testing.T) {
    cmd := New("bar", nil)
    assert.Equal(t, "bar", cmd.Description)
}

func TestNew_NoSubs(t *testing.T) {
    cmd := New("bar", nil)
    assert.Equal(t, 0, len(cmd.Subs))
}

func TestNewSub_ParentSubs(t *testing.T) {
    cmd := New("bar", nil)
    cmd.NewSub("car", "dar", "dar", nil)
    assert.Equal(t, 1, len(cmd.Subs))
}

func TestNewSub_SetName(t *testing.T) {
    cmd := New("bar", nil)
    sub := cmd.NewSub("car", "dar", "dar", nil)
    assert.Equal(t, "car", sub.Name)
}

func TestNewSub_SetDescription(t *testing.T) {
    cmd := New("bar", nil)
    sub := cmd.NewSub("car", "dar", "dar", nil)
    assert.Equal(t, "dar", sub.Description)
}

func TestNewSub_NoSubs(t *testing.T) {
    cmd := New("bar", nil)
    sub := cmd.NewSub("foo", "bar", "bar", nil)
    assert.Equal(t, 0, len(sub.Subs))
}

func TestHasCallback_NoCallback(t *testing.T) {
    cmd := New("bar", nil)
    assert.False(t, cmd.HasCallback())
}

type TestHasCallback_WithCallbackCmd struct {}

func (opts *TestHasCallback_WithCallbackCmd) Run() error {
    return nil
}

func TestHasCallback_WithCallback(t *testing.T) {
    cmd := New("bar", &TestHasCallback_WithCallbackCmd{})
    assert.True(t, cmd.HasCallback())
}

func TestHasSubs_NoSubs(t *testing.T) {
    cmd := New("bar", nil)
    assert.False(t, cmd.HasSubs())
}

func TestHasSubs_WithSubs(t *testing.T) {
    cmd := New("bar", nil)
    cmd.NewSub("foo", "bar", "bar", nil)
    assert.True(t, cmd.HasSubs())
}

func TestHasSub_NoSubs(t *testing.T) {
    cmd := New("bar", nil)
    assert.False(t, cmd.HasSub("bar"))
}

func TestHasSub_NoChild(t *testing.T) {
    cmd := New("bar", nil)
    cmd.NewSub("foo", "bar", "bar", nil)
    assert.False(t, cmd.HasSub("bar"))
}

func TestHasSub_WithChild(t *testing.T) {
    cmd := New("bar", nil)
    cmd.NewSub("foo", "bar", "bar", nil)
    assert.True(t, cmd.HasSub("foo"))
}

var TestRun_ErrorCalled bool

type TestRun_ErrorCmd struct {
    t *testing.T
    Verbose bool `default:"Nein" short:"v"`
}

func (opts *TestRun_ErrorCmd) Run() error {
    TestRun_ErrorCalled = true
    return nil
}

func TestRun_Error(t *testing.T) {
    assert.False(t, TestRun_ErrorCalled)
    cmd := New("bar", &TestRun_ErrorCmd{t: t})
    err := cmd.Run([]string{"-v"})
    assert.NotNil(t, err)
    assert.False(t, TestRun_ErrorCalled)
}

var TestRun_ArgsCalled bool

type TestRun_ArgsCmd struct {
    t *testing.T
    Verbose bool `short:"v" description:"Show verbose debug information"`
}

func (opts *TestRun_ArgsCmd) Run() error {
    TestRun_ArgsCalled = true
    assert.True(opts.t, opts.Verbose)
    return nil
}

func TestRun_WithArgs(t *testing.T) {
    assert.False(t, TestRun_ArgsCalled)
    cmd := New("bar", &TestRun_ArgsCmd{t: t})
    err := cmd.Run([]string{"-v"})
    assert.Nil(t, err)
    assert.True(t, TestRun_ArgsCalled)
}

var TestRun_LeftoverArgsCalled bool

type TestRun_LeftoverArgsCmd struct {
    t *testing.T
    Args []string `positional:"true"`
    Verbose bool `short:"v" description:"Show verbose debug information"`
}

func (opts *TestRun_LeftoverArgsCmd) Run() error {
    TestRun_LeftoverArgsCalled = true
    assert.Equal(opts.t, []string{"foo", "bar"}, opts.Args)
    return nil
}

func TestRun_LeftoverArgs(t *testing.T) {
    assert.False(t, TestRun_LeftoverArgsCalled)
    cmd := New("bar", &TestRun_LeftoverArgsCmd{t: t})
    err := cmd.Run([]string{"-v", "foo", "bar"})
    assert.Nil(t, err)
    assert.True(t, TestRun_LeftoverArgsCalled)
}

func TestRun_NoCallback(t *testing.T) {
    cmd := New("bar", nil)
    err := cmd.Run(nil)
    assert.NotNil(t, err)
}

var TestSubRun_LeftoverArgsCalled bool

type TestSubRun_LeftoverArgsCmd struct {
    t *testing.T
    Args []string `positional:"true"`
}

func (opts *TestSubRun_LeftoverArgsCmd) Run() error {
    TestSubRun_LeftoverArgsCalled = true
    assert.Equal(opts.t, []string{"car", "dar", "far"}, opts.Args)
    return nil
}

func TestSubRun_LeftoverArgs(t *testing.T) {
    assert.False(t, TestSubRun_LeftoverArgsCalled)
    cmd := New("foo", nil)
    cmd.NewSub("bar", "dar", "dar", &TestSubRun_LeftoverArgsCmd{t: t})
    err := cmd.Run([]string{"bar", "car", "dar", "far"})
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
    cmd := New("foo", nil)
    cmd.NewSub("bar", "dar", "dar", &TestSubRun_CallCallbackCmd{t: t})
    err := cmd.Run([]string{"bar"})
    assert.Nil(t, err)
    assert.True(t, TestSubRun_CallCallbackCalled)
}

func TestSubRun_NoCallback(t *testing.T) {
    cmd := New("foo", nil)
    cmd.NewSub("bar", "dar", "dar", nil)
    err := cmd.Run([]string{"bar"})
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
    cmd := New("foo", nil)
    sub := cmd.NewSub("bar", "dar", "dar", nil)
    sub.NewSub("far", "car", "car", &TestNestedSubRun_CallCallbackCmd{t: t})
    err := cmd.Run([]string{"bar", "far"})
    assert.Nil(t, err)
    assert.True(t, TestNestedSubRun_CallCallbackCalled)
}
