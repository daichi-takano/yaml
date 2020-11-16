package main                                           

import (                                               
    "fmt"                                              
    "github.com/go-yaml/yaml"                          
    "io/ioutil"                                        
)

// マッピングする構造体です。
type Test struct {  
    First  string   
    Second string   
}                                                          

func main() {                                          
    buf, err := ioutil.ReadFile("./test.yaml")               
    if err != nil {                                    
        fmt.Println(err)                               
        return                                         
    }       

    data, err := ReadOnStruct(buf)                    
    if err != nil {                                
        fmt.Println(err)                           
        return                                     
    }     

    // hello を hi に書き換えて保存します。                                         
    data[0].First = "hi"             
    err = WriteOnFile(fileName, data)
    if err != nil {                  
        fmt.Println(err)             
        return                       
    }                                                        
}                                                      

// yaml形式の[]byteを渡すと[]Testに変換してくれる関数です。
func ReadOnStruct(fileBuffer []byte) ([]Test, error) {
    data := make([]Test, 20) 
    // []map[string]string のときと使う関数は同じです。
    // いい感じにマッピングしてくれます。                  
    err := yaml.Unmarshal(fileBuffer, &data)          
    if err != nil {                                   
        fmt.Println(err)                              
        return nil, err                               
    }                                                 
    return data, nil                                  
}

// ファイル名とデータをを渡すとyamlファイルに保存してくれる関数です。
func WriteOnFile(fileName string, data interface{}) error { 
    // ここでデータを []byte に変換しています。
    buf, err := yaml.Marshal(data)                          
    if err != nil {                                         
        return err                                          
    }                                      
    // []byte をファイルに上書きしています。                 
    err = ioutil.WriteFile(fileName, buf, os.ModeExclusive) 
    if err != nil {                                         
        return err                                          
    }                                                       
    return nil                                              
}                 }                  