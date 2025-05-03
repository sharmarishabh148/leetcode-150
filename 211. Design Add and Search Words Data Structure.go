type TrieNode struct{
    Children  map[rune]*TrieNode // or [26]*TrieNode
    IsEnd bool
}

type WordDictionary struct {
    root *TrieNode
}


func Constructor() WordDictionary {
    return WordDictionary{
        root: &TrieNode{
            Children: make(map[rune]*TrieNode),
        },
    }
}


func (this *WordDictionary) AddWord(word string)  {
    node := this.root
    for _, ch := range word{ // returns rune
        if node.Children[ch] == nil{
            node.Children[ch] = &TrieNode{Children: make(map[rune]*TrieNode)}
        }
        node = node.Children[ch]
    }
    node.IsEnd = true
}


func (this *WordDictionary) Search(word string) bool {
    return dfs(this.root, word, 0)
}


func dfs(node *TrieNode, word string, index int) bool {
    if node == nil{
        return false
    }
    if index==len(word){
        return node.IsEnd
    }
    ch := rune(word[index])
    if ch =='.'{
        for _, child := range node.Children{
            if dfs(child, word, index+1){
                return true
            }
        }
        return false
    }
    return dfs(node.Children[ch], word, index+1)
}
/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
