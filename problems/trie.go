package problems 

type Trie struct {
	wordEnd bool
	child [26] *Trie
}


func TrieConstructor() Trie {
	newTrie := Trie{}
	newTrie.wordEnd = false
	
	for i:=0 ; i<len(newTrie.child); i++ {
		newTrie.child[i] = nil
	}
	return newTrie
}


func (this *Trie) TrieInsert(word string)  {
	current := this
	for _ , char := range word {
		if current.child[char - 'a'] == nil {
			newTrie := TrieConstructor()
			current.child[char-'a'] = &newTrie
		}

		current = current.child[char - 'a']
	}
	current.wordEnd = true
}


func (this *Trie) TrieSearch(word string) bool {
	current := this

	for _, c := range word {
		if current.child[c-'a'] == nil {
			return false
		}
		current = current.child[c-'a']
	}
	return current.wordEnd
}


func (this *Trie) TrieStartsWith(prefix string) bool {
	current := this

	for _, c := range prefix {
		if current.child[c-'a'] == nil {
			return false
		}
		current = current.child[c-'a']
	}
	return true;
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
