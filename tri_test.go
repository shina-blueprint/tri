package tri

import "testing"

func TestIf(t *testing.T) {
	result := If(true, "true value", "false value")
	if result != "true value" {
		t.Errorf("Expected 'true value', got %s", result)
	}

	result = If(false, "true value", "false value")
	if result != "false value" {
		t.Errorf("Expected 'false value', got %s", result)
	}
}

func TestIfExec(t *testing.T) {
	result := IfExec(true, func() string { return "true func" }, func() string { return "false func" })
	if result != "true func" {
		t.Errorf("Expected 'true func', got %s", result)
	}

	result = IfExec(false, func() string { return "true func" }, func() string { return "false func" })
	if result != "false func" {
		t.Errorf("Expected 'false func', got %s", result)
	}
}

func TestIfExecOrVal(t *testing.T) {
	result := IfExecOrVal(true, func() string { return "true func" }, "false value")
	if result != "true func" {
		t.Errorf("Expected 'true func', got %s", result)
	}

	result = IfExecOrVal(false, func() string { return "true func" }, "false value")
	if result != "false value" {
		t.Errorf("Expected 'false value', got %s", result)
	}
}

func TestIfValOrExec(t *testing.T) {
	result := IfValOrExec(true, "true value", func() string { return "false func" })
	if result != "true value" {
		t.Errorf("Expected 'true value', got %s", result)
	}

	result = IfValOrExec(false, "true value", func() string { return "false func" })
	if result != "false func" {
		t.Errorf("Expected 'false func', got %s", result)
	}
}
