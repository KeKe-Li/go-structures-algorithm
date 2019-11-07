### [1. Two Sum](https://leetcode.com/problems/two-sum/)

#### 题目两数之和

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```text
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

题意:

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
```text

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

```

#### 解题思路

```go
a + b = target
```

也可以看成是

```go
a = target - b
```

在`map[整数]整数的序号`中，可以查询到a的序号。这样就不用嵌套两个for循环了。


#### 测试用例

Test的测试用例有四种形式：
 
* TestXxxx(t testing.T) 基本测试用例

单元测试中，传递给测试函数的参数是 *testing.T 类型。它用于管理测试状态并支持格式化测试日志。
测试日志会在执行测试的过程中不断累积，并在测试完成时转储至标准输出。

```go
func TestTwoSum(t *testing.T){
	nums := []int{2,7,11,15}
	targets := 9
	t.Log(twoSum(nums,targets))
}
```
 
* BenchmarkXxxx(b testing.B) 压力测试的测试用例

B 是传递给基准测试函数的一种类型，它用于管理基准测试的计时行为，并指示应该迭代地运行测试多少次。
```go
func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
```
运行是采用:
```bash
> go test -bench=.
```

跟单元测试一样，基准测试会在执行的过程中积累日志，并在测试完毕时将日志转储到标准错误。但跟单元测试不一样的是，为了避免基准测试的结果受到日志打印操作的影响，基准测试总是会把日志打印出来。

B 类型中的报告方法使用方式和 T 类型是一样的，一般来说，基准测试中也不需要使用，毕竟主要是测性能。这里我们对 B 类型中其他的一些方法进行讲解。


* Example_Xxx() 测试控制台输出的例子

Example的例子:（Example需要在最后用注释的方式确认控制台输出和预期是不是一致的） 
```go
func Example_GetScore() {
    score := getScore(100, 100, 100, 2.1)
    fmt.Println(score)
    // Output:
    // 31.1
}
``` 
 
* TestMain(m *testing.M)  测试Main函数  

有时，测试还需要控制在主线程上运行的代码。为了支持这些需求，testing 提供了 TestMain 函数:
```go
func TestMain(m *testing.M)
```

如果测试文件中包含该函数，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。TestMain 运行在主 goroutine 中, 可以在调用 m.Run 前后做任何设置和拆卸。
注意，在 TestMain 函数的最后，应该使用 m.Run 的返回值作为参数调用 os.Exit。




#### Test的变量

gotest的变量有这些：

* test.short : 一个快速测试的标记，在测试用例中可以使用testing.Short()来绕开一些测试
* test.outputdir : 输出目录
* test.coverprofile : 测试覆盖率参数，指定输出文件
* test.run : 指定正则来运行某个/某些测试用例
* test.memprofile : 内存分析参数，指定输出文件
* test.memprofilerate : 内存分析参数，内存分析的抽样率
* test.cpuprofile : cpu分析输出参数，为空则不做cpu分析
* test.blockprofile : 阻塞事件的分析参数，指定输出文件
* test.blockprofilerate : 阻塞事件的分析参数，指定抽样频率
* test.timeout : 超时时间
* test.cpu : 指定cpu数量
* test.parallel : 指定运行测试用例的并行数

#### Test包内的结构

* B : 压力测试
* BenchmarkResult : 压力测试结果
* Cover : 代码覆盖率相关结构体
* CoverBlock : 代码覆盖率相关结构体
* InternalBenchmark : 内部使用的结构
* InternalExample : 内部使用的结构
* InternalTest : 内部使用的结构
* M : main测试使用的结构
* PB : Parallel benchmarks 并行测试使用结果
* T : 普通测试用例
* TB : 测试用例的接口

#### Test的通用方法

T结构内部是继承自common结构，common结构提供集中方法，是我们经常会用到的：

当我们遇到一个断言错误的时候，我们就会判断这个测试用例失败，就会使用到：
```go
Fail  : case失败，测试用例继续
FailedNow : case失败，测试用例中断
```

当我们遇到一个断言错误，只希望跳过这个错误，但是不希望标示测试用例失败，会使用到：

```go
SkipNow : case跳过，测试用例不继续
```
当我们只希望在一个地方打印出信息，我们会用到:

```go
Log : 输出信息
Logf : 输出有format的信息
```

当我们希望跳过这个用例，并且打印出信息:

```go
Skip : Log + SkipNow
Skipf : Logf + SkipNow
```
当我们希望断言失败的时候，测试用例失败，打印出必要的信息，但是测试用例继续：

```go
Error : Log + Fail
Errorf : Logf + Fail
```

当我们希望断言失败的时候，测试用例失败，打印出必要的信息，测试用例中断：

```go

Fatal : Log + FailNow
Fatalf : Logf + FailNow
```