type TrieNode struct{
    Children  map[rune]*TrieNode // or [26]*TrieNode
    IsEnd bool
}

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    return Trie{
        root: &TrieNode{
            Children: make(map[rune]*TrieNode),
        },
    }
}


func (this *Trie) Insert(word string)  {
    node := this.root
    for _, ch := range word{ // returns rune
        if node.Children[ch] == nil{
            node.Children[ch] = &TrieNode{Children: make(map[rune]*TrieNode)}
        }
        node = node.Children[ch]
    }
    node.IsEnd = true
}


func (this *Trie) Search(word string) bool {
    node := this.root
    for _, ch := range word{
        if node.Children[ch] ==nil{
            return false
        }
        node = node.Children[ch]
    }
    return node.IsEnd
}


func (this *Trie) StartsWith(prefix string) bool {
     node := this.root
    for _, ch := range prefix{
        if node.Children[ch] ==nil{
            return false
        }
        node = node.Children[ch]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
