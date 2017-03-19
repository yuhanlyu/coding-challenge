/* Implement a trie with insert, search, and startsWith methods.
 * Time Complexity: O(n)
 * Space Complexity: O(n)
 */

/**
 * @constructor
 * Initialize your data structure here.
 */
var TrieNode = function() {
    this.map = new Map();
};

var Trie = function() {
    this.root = new TrieNode();
};

/**
 * @param {string} word
 * @return {void}
 * Inserts a word into the trie.
 */
Trie.prototype.insert = function(word) {
    var node = this.root;
    for (var c of word) {
        if (!node.map.has(c))
            node.map.set(c, new TrieNode());
        node = node.map.get(c);
    }
    node.map.set(null, null);
};

/**
 * @param {string} word
 * @return {boolean}
 * Returns if the word is in the trie.
 */
Trie.prototype.search = function(word) {
    var node = this.root;
    for (var c of word) {
        if (node.map.has(c))
            node = node.map.get(c);
        else
            return false;
    }
    return node.map.has(null);
};

/**
 * @param {string} prefix
 * @return {boolean}
 * Returns if there is any word in the trie
 * that starts with the given prefix.
 */
Trie.prototype.startsWith = function(prefix) {
    var node = this.root;
    for (var c of prefix) {
        if (node.map.has(c))
            node = node.map.get(c);
        else
            return false;
    }
    return true;
};

/**
 * Your Trie object will be instantiated and called as such:
 * var trie = new Trie();
 * trie.insert("somestring");
 * trie.search("key");
 */
