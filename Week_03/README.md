### 递归

递归适用的情况，是把问题能分解成更小规模的、具有和原来问题相同解法的情况。关键是通过递归缩小问题规模。

递归模板：

```go
func recursion(level int,param int){
    //terminator
    if level > MAX_LEVEL{
        //process result
        return
    }
    
    //process current logic
    process(level,param)
    
    //drill down
    recursion(level+1, newParam)
    
    //restore current status
}
```

### 回溯

回溯法主要是采用试错的思路，尝试分步解决问题。在分步解决问题的过程中，通过尝试发现现有的分步答案不能得到有效的正确的解答时，就取消上一步甚至上几步的计算，在通过其他可能的分步解答尝试寻找答案。

不断的在每一层试结果是否成立，本质还是递归（常适用问题：八皇后、数组）

回溯通常用最简单的递归方法实现，最后可能出现存在正确答案或者尝试所有有可能的分步方法后发现没有答案

### 分治

分治是将求解的原来问题划分成k个较小规模的子问题，对这k个子问题分别求解。如果子问题的规模仍然不够小，则在将每个子问题划分成k个规模更小的子问题，如此分解下去，直到问题规模足够小，求出解为止，再将子问题的解合并成更大规模的问题的解，自底向上逐步求出原问题的解。

分治模板：

```go
func divide_conquer(problem,param1,param2,...){
    //recursion terminator
    if problem == nil{
        return
    }
    
    //prepare data
    data := prepare_data(problem)
    subproblems = split_problem(problem,data)
    
    //conquer subproblems
    subresult1 = divide_conquer(subproblems[0],p1,p1,...)
    subresult2 = divide_conquer(subproblems[2],p1,p1,...)
    subresult3 = divide_conquer(subproblems[3],p1,p1,...)    
    
    //process and generate the final result
    result = process_result(subresult1,subresult2,subresult3,...)
    
    //revert the current level states
}
```

