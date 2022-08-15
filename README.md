# batch_raname
a tool to batch rename files in a certain directory
## build
0. git clone https://github.com/shentupenghui/batch_raname.git
1. `go build -o batch-rename main.go`

*in addition, one more step for linux、mac*
2. `chmod +x batch-rename`

## usage
### delete some str in filename for all files in a certain directory
`batch-rename test_dir delete_str`
### replace some str in filename for all files in a certain directory
`batch-rename test_dir replace_str new_str`



