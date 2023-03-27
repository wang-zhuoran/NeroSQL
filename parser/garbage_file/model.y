// %{
// package main

// import (
//     "fmt"
//     "strings"
// )

// type Model struct {
//     Name     string   // 模型名称
//     Args     []string // 参数列表
//     Table    string   // 数据表名
//     Function string   // 模型函数名称
// }

// var currentModel *Model
// var models []*Model
// %}

// %start create_model

// %token <str> ID
// %token <str> STRING
// %token COMMA

// %%

// create_model:
//     CREATE MODEL ID AS select_stmt ';' { 
//         currentModel.Name = $3
//     }

// select_stmt:
//     SELECT model_call AS pred_output FROM ID { 
//         currentModel.Function = $2
//         currentModel.Table = $6
//         models = append(models, currentModel)
//         currentModel = nil
//     }



// /*
// $1实际上就是一个ID Token，
// 它表示调用的函数名。
// 而$3则是一个arg_list，它表示函数调用时传入的参数列表（可能为空或者包含多个标识符）。
// 因此，$2代表了左括号 "("，而$4代表了右括号 ")"。
// */

// model_call:
//     ID '(' arg_list ')' { $$ = $1 + "(" + strings.Join($3, ", ") + ")" }
// ;

// arg_list:
//     /* empty */ { $$ = []string{} } 
//     | arg_list_single { $$ = []string{$1} } //单一参数
//     | arg_list ',' arg_list_single { $$ = append($1.([]string), $3) } //多个参数
// ;

// arg_list_single:
//     {$$ = $1}
// ;

// %%

// func parse(input string) ([]*Model, error) {
//     currentModel = &Model{}
//     models = []*Model{}

//     lexer := lex(input)
//     if yyParse(NewParser(lexer)) != 0 {
//         return nil, fmt.Errorf("parse error")
//     }

//     return models, nil
// }

// func (p *parserImpl) Error(s string) {
//     panic(fmt.Sprintf("parse error: %s", s))
// }
