package assert

import (
    "reflect"
    "strings"
    "testing"
)

func Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("expected %v, got %v", expected, actual)
        return false
    }
    return true
}

func NoError(t *testing.T, err error, msgAndArgs ...interface{}) bool {
    if err != nil {
        t.Errorf("unexpected error: %v", err)
        return false
    }
    return true
}

func Len(t *testing.T, object interface{}, length int, msgAndArgs ...interface{}) bool {
    v := reflect.ValueOf(object)
    if v.Len() != length {
        t.Errorf("expected length %d, got %d", length, v.Len())
        return false
    }
    return true
}

func NotNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
    if isNil(object) {
        t.Errorf("expected non-nil")
        return false
    }
    return true
}

func Nil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
    if !isNil(object) {
        t.Errorf("expected nil, got %#v", object)
        return false
    }
    return true
}

func Error(t *testing.T, err error, msgAndArgs ...interface{}) bool {
    if err == nil {
        t.Errorf("expected error")
        return false
    }
    return true
}

func True(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
    if !value {
        t.Errorf("expected true")
        return false
    }
    return true
}

func Contains(t *testing.T, s, substr string, msgAndArgs ...interface{}) bool {
    if !strings.Contains(s, substr) {
        t.Errorf("%q does not contain %q", s, substr)
        return false
    }
    return true
}

func isNil(i interface{}) bool {
    if i == nil {
        return true
    }
    v := reflect.ValueOf(i)
    switch v.Kind() {
    case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
        return v.IsNil()
    default:
        return false
    }
}
