# chrononz
chrononz takes approximate (minimum) timestamp of a Golang ELF binary.

More info in blogpost: [RU](https://www.orderofsixangles.com/ru/2022/07/09/goelf-time-ru.html)/[EN](https://www.orderofsixangles.com/en/2022/07/09/goelf-time-en.html)

Algorithm:

1. Get a list of 3rd party dependencies (using github.com/goretk/gore)
2. Get a version of each dependency
3. If it is github dependency, then use Github API to fetch version date
4. Make a list {dependency:date}
5. Take the maximum date

