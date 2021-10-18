class Solution {
public:
    string removeDuplicateLetters(string s) {
        // 用map, set, 不用数组的话, 可以提高泛用性，就是多费点空间
		//count用于记录每个字母在s中出现的次数
        unordered_map<char, int> count;
		//利用hashset记录每个字母是否存在栈内
        unordered_set<char> us;
		//单调递减栈(栈底->栈顶:元素值递增)
        stack<char> st;
		//初始化count
        for (int i = 0; i < s.size(); ++i) ++count[s[i]];
		//单调栈的一般步骤，对s字符串进行遍历
        for (int i = 0; i < s.size(); ++i) {
            char c = s[i];
            --count[c];//这一步比较巧，每遍历一次都减1，表示当前字母已经被遍历过一次
			//当在set us里面没有找到当前字母时进入循环
            if (us.find(c) == us.end()) {
				//单调栈模板: 注意LZ大佬这么写的话(st.top() > c)就是>而不是>= 
				//因为如果栈顶元素和当前遍历的字母相同，直接就被跳过了, 无需进入if
                while (!st.empty() && st.top() > c && count[st.top()] != 0) {
					//这个地方为什么是!=0 是因为之前每个字母都会-1，也就是遍历过1次就会-1
					//而如果一个字母只出现过一次的话就==0了
					//利用单调栈维护所要的字典序最小这一目的，也就是尽可能地让位于前面的元素值小
                    us.erase(st.top());//在set中抹除掉栈顶元素
                    st.pop();
                }
                us.emplace(c);//将c加入set
                st.push(c);
            }
        }
        string ans = "";
		//得到最终的ans
        while (!st.empty()) {
            ans.push_back(st.top());
            st.pop();
        }
        reverse(ans.begin(), ans.end());//由于stack是倒序的，需要颠倒过来
        return ans;
    }
};
