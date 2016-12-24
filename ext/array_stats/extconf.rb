require 'mkmf'
find_executable('go')
create_makefile('array_stats/array_stats')
exec 'go build -buildmode=c-shared -o array_stats.so ../../../../ext/array_stats/array_stats.go'
