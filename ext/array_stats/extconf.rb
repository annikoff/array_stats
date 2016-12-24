require 'mkmf'
find_executable('go')
%x{go build -buildmode=c-shared -o libfastpercentile.so fast_percentile.go}
create_makefile('array_stats/array_stats')
