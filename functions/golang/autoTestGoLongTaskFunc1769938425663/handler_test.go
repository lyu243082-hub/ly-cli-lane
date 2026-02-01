package autoTestGoLongTaskFunc1769938425663

import (
    "context"
    "testing"
)

func TestHandler(t *testing.T) {
    result, err := Handler(context.Background(), &Params{})
    if err != nil {
        t.Fatalf("失败: %v", err)
    }
    t.Logf("成功: %+v", result)
}
