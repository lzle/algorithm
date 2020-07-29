# 二分查找

前提：`排序好的数组。`

方法：`每次取中间值进行比较，高了向下取，低了向上取。`

二分查找易出错的几点：

* `循环退出条件，注意是 low<=high，而不是 low<high。`

* `mid的取值  low+(high-low)/2。`

* `low和high的更新，low=mid+1，high=mid-1。不能写成low=mid，high=mid，容易死循环。`

二分查找依赖于`有序数组`这种数据结构，不适合数据量太小、太大的情况。时间复杂度`O(logn)。`

## 变形问题

#### 查找第一个值等于给定值的元素

`注意只要一个元素时 s=[1] v=1`

``` Go
func binarySearch(s []int, v int) int{
    low := 0
    high := len(s) - 1
    mid := 0
   
    for low <= high {
        mid = low + (high-low)>>2
        if s[mid] > v {
    	    high = mid -1
     	} else if s[mid] < v {
            low = mid + 1
        } else {
            if mid == 0 || s[mid-1] != v {
               return mid
            }
            high = mid -1
        }
    }
    return -1
}
```

#### 查找最后一个值等于给定值的元素

`注意最后一个元素为查找值时 s=[1,2,3],v=3`

````
func binarySearch(s []int, v int) int {
    low := 0
    high := len(s) - 1
    mid := 0

    for low <= high {
    	mid = low + (high-low)>>2
    	if s[mid] > v {
    		high = mid - 1
    	} else if s[mid] < v {
    		low = mid + 1
    	} else {
    		if mid+1 > len(s)-1 || s[mid+1] != v {
    			return mid
    		}
    		low = mid + 1
    	}
    }
    return -1
}
````

#### 查找第一个大于等于给定值的元素

`思路与第一个值等于给定值一致`

``` Go
func binarySearch(s []int, v int) int{
    low := 0
    high := len(s) - 1
    mid := 0
   
    for low <= high {
        mid = low + (high-low)>>2
        if s[mid] >= v {
    	    if mid == 0 || s[mid-1] < v {
    	    	return mid
            }
            high = mid -1
     	} else if s[mid] < v {
            low = mid + 1
        } else {
            if mid == 0 || s[mid-1] != v {
               return mid
            }
            high = mid -1
        }
    }
    return -1
}
```

#### 查找最后一个小于等于给定值的元素

``` Go
func binarySearch(s []int, v int) int {
    low := 0
    high := len(s) - 1
    mid := 0
    for low <= high {
    	mid = low + (high-low)>>2
    	if s[mid] > v {
    	    high = mid - 1
    	} else if s[mid] <= v {
    	    if mid == len(s)-1 || s[mid+1] > v {
    	    	return mid
    	    }
    	    low = mid + 1
    	}   
    }
    return -1
}
```

## 小结

凡是用二分查找能解决的，绝大部分我们更倾向于用`散列表`或者`二叉查找树`。即便是二分查找在内存使用上更节省，但是毕竟内存如此紧缺的情况并不多。

二分查找更适合用在`“近似”查找问题`，在这类问题上，二分查找的优势更加明显。上面几种变体问题，用其他数据结构，比如散列表、二叉树，就比较难实现了。

## 思考 🤔

1）如何查询 IP 归属地？

2）数据使用链表存储，二分查找的时间复杂多少？