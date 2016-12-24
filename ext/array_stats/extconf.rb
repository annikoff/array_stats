require 'mkmf'
find_executable('go')
%x(go build -buildmode=c-shared -o libfastpercentile.so fast_percentile.go)
create_makefile('ext/array_stats/fast_percentile')