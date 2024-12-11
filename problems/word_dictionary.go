package problems

type WordDictionary struct {
	wordEnd bool
	child [27]*WordDictionary
}


func Constructor() WordDictionary {
	dictionary :=  WordDictionary{}
	dictionary.wordEnd  = false
	for i:=0; i<27; i++ {
		dictionary.child[i] = nil
	}
	return dictionary
}


func (this *WordDictionary) AddWord(word string)  {
	current := this
	for _, c := range word {
		if current.child[c - 'a'] == nil {
			newDict := Constructor();
			current.child[c - 'a'] = &newDict
		}
		current = current.child[c-'a']
	}
	current.wordEnd = true
}


func (this *WordDictionary) Search(word string) bool {
	// base case 
	var dfs func(node *WordDictionary, idx int) bool
	dfs = func(node *WordDictionary, idx int) bool {
		if idx == len(word) {
			return node.wordEnd
		}

		c := word[idx]

		if c == '.' {
			for i := 0; i < 26; i++ {
				if node.child[i] != nil && dfs(node.child[i], idx+1) {
					return true
				}
			}
			return false
		} else {
			if node.child[c - 'a'] == nil {
				return false
			}
			return dfs(node.child[c-'a'], idx+1)
		}
	}
	
	return dfs(this, 0)
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */