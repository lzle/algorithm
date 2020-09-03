# Trie

像 Google、百度这样的搜索引擎，它们的关键词提示功能非常全面和精准，底层最基本的原理就是今天要讲的这种数据结构：Trie 树。

Trie 树，也叫“字典树”。它是一种专门处理字符串匹配的数据结构，用来解决在一组字符串集合中快速查找某个字符串的问题。

针对字符串 how，hi，her，hello，so，see ，构造出来的就是下面这个图中的样子。

<img src="https://static001.geekbang.org/resource/image/28/32/280fbc0bfdef8380fcb632af39e84b32.jpg" width=500>

Trie 树的本质，就是利用字符串之间的公共前缀，将重复的前缀合并在一起。

Trie 树主要有两个操作：构建、查询。

### 存储

最经典的存储方式是借助散列表的思想，通过一个下标与字符一一映射的数组，来存储子节点的指针。

<img src="https://static001.geekbang.org/resource/image/f5/35/f5a4a9cb7f0fe9dcfbf29eb1e5da6d35.jpg" width=500>

思想

``` java
public class Trie {
  private TrieNode root = new TrieNode('/'); // 存储无意义字符

  // 往Trie树中插入一个字符串
  public void insert(char[] text) {
    TrieNode p = root;
    for (int i = 0; i < text.length; ++i) {
      int index = text[i] - 'a';
      if (p.children[index] == null) {
        TrieNode newNode = new TrieNode(text[i]);
        p.children[index] = newNode;
      }
      p = p.children[index];
    }
    p.isEndingChar = true;
  }

  // 在Trie树中查找一个字符串
  public boolean find(char[] pattern) {
    TrieNode p = root;
    for (int i = 0; i < pattern.length; ++i) {
      int index = pattern[i] - 'a';
      if (p.children[index] == null) {
        return false; // 不存在pattern
      }
      p = p.children[index];
    }
    if (p.isEndingChar == false) return false; // 不能完全匹配，只是前缀
    else return true; // 找到pattern
  }

  public class TrieNode {
    public char data;
    public TrieNode[] children = new TrieNode[26];
    public boolean isEndingChar = false;
    public TrieNode(char data) {
      this.data = data;
    }
  }
}
```

构建 Trie 树的过程，需要扫描所有的字符串，时间复杂度是 O(n)（n 表示所有字符串的长度和）。
查询时，效率很高，如果要查询的字符串长度是 k，那我们只需要比对大约 k 个节点，就能完成查询操作。

### 很耗内存吗？

Trie 树是一种非常独特的、高效的字符串匹配方法。也非常的耗内存，用的是一种空间换时间的思路。

用数组来存储一个节点的子节点的指针，每个数组元素要存储一个 8 字节指针（或者是 4 字节，这个大小跟 CPU、操作系统、编译器等有关）。
如果数组长度为 26，每个元素是 8 字节，那每个节点就会额外需要 26*8=208 个字节。而且这还是只包含 26 个字符的情况。

可以稍微牺牲一点查询的效率，将每个节点中的数组换成其他数据结构，比如有序数组、跳表、散列表、红黑树等。


### Trie 树与散列表、红黑树的比较

字符串的匹配问题，笼统上讲，其实就是数据的查找问题。对于支持动态数据高效操作的数据结构，有很多，比如散列表、红黑树、跳表等等。

Trie 树对要处理的字符串有及其严苛的要求。

第一，字符串中包含的字符集不能太大。我们前面讲到，如果字符集太大，那存储空间可能就会浪费很多。即便可以优化，但也要付出牺牲查询、插入效率的代价。

第二，要求字符串的前缀重合比较多，不然空间消耗会变大很多。

第三，如果要用 Trie 树解决问题，那我们就要自己从零开始实现一个 Trie 树，还要保证没有 bug，这个在工程上是将简单问题复杂化，除非必须，一般不建议这样做。

第四，我们知道，通过指针串起来的数据块是不连续的，而 Trie 树中用到了指针，所以，对缓存并不友好，性能上会打个折扣。

实际上，Trie 树只是不适合`精确匹配查找`，这种问题更适合用散列表或者红黑树来解决。Trie 树比较适合的是`查找前缀匹配的字符串`，也就是类似开篇问题的那种场景。
