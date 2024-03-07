package main

import (
    "testing"
    "os"
    "time"
    "fmt"
    "bufio"

)


func (c *Engine) GetFileContent(f *os.File) []string {
    c.mu.Lock()
    defer c.mu.Unlock()

    _, err := f.Seek(0, 0)
    if err != nil {
        fmt.Println(err)
        return []string{}
    }

    scanner := bufio.NewScanner(f)

    var content []string
    for scanner.Scan() {
        line := scanner.Text()
        content = append(content, line)
    }

    return content
}
 
func TestEngine_Compact(t *testing.T) {
    os.Remove("data.txt")
    v1 := "latestvalue1"
    v2 := "latestvalue2"
    e, _ := NewEngine()
    e.Set("key1", "value1")
    e.Set("key2", "value2")
    e.Set("key1", v1)
    e.Set("key2", v2)
    e.Set("key3", "value3")

    go e.CompactFile()

    time.Sleep((Seconds + 3) * time.Second)
    if len(e.GetFileContent(e.file)) != 3 {
        t.Errorf("Expected %d, but got %d", 3, len(e.GetFileContent(e.file)))
    }

}
/*
func TestEngine_Restore(t *testing.T) {
    os.Remove("data.txt")
    e, _ := NewEngine()

    e.Set("key1_restore", "value1")
    e.Set("key2_restore", "value2")

    e.Close()

    e, _ = NewEngine()
    e.Restore()
    k, _ := e.Get("key1_restore")

    if k != "value1" {
        t.Errorf("Expected %s, but got %s", "value1", k)
    }
}

*/
/*
func TestEngine_DeleteKey(t *testing.T) {
    os.Remove("data.txt")
    os.Remove("remove.txt")
    e, _ := NewEngine()

    e.Set("key1_delete", "value1")
    e.Set("key2_delete", "value2")

    err := e.Delete("key1_delete")
    if err != nil {
        panic(err)
    }

    k, _ := e.Get("key1_delete")

    if k != "" {
        t.Errorf("Expected %s, but got %s", "", k)
    }

    if len(e.GetFileContent(e.fileDelete)) != 1 {
        t.Errorf("Expected %d, but got %d", 1, len(e.GetFileContent(e.file)))
    }
}
*/
/*
func TestEngine_DeleteKeyFromFile(t *testing.T) {
    os.Remove("data.txt")
    //os.Remove("delete.txt")
    e, _ := NewEngine()

    e.Set("key1_delete", "value1")
    e.Set("key22_delete", "value2")
    e.Set("key33_delete", "value3")

    e.deleteKeyFromFile([]string{"key22_delete", "key33_delete"})

    if len(e.GetFileContent(e.file)) != 1 {
        t.Errorf("Expected %d, but got %d", 1, len(e.GetFileContent(e.file)))
    }
}

*/


var cfg = Config{
    FileData:   "data.txt",
    FileRemove: "delete.txt",
}

func Test_SetGetKeyValue(t *testing.T) {
    e, _ := NewEngine(cfg)
    e.Set("test", "data")
    e.Set("foo", "bar")
    value, err := e.Get("foo")
    if err != nil {
        t.Error(err)
    }
    if value != "bar" {
        t.Error("value should be bar")
    }

    _, err = e.Get("notfound")
    if err == nil {
        t.Error("should return error")
    }
}