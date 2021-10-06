//在一个数组nums中 找到一个三元组(a,b,c) 使得 nums[b] > nums[a] 并且 nums[b] > nums[c], c > b > a, 多个答案输出字典序小的
//尽可能使a小, 字典序a优先级最高然后是b最后是c

//两步操作: 第一步是利用一个单调递减栈实现对b的记录，找到每个b对应的离它最近的小于它的数的索引，也就是c
//第二步操作是利用二分查找找到索引最小的a并且同时满足nums[a] < nums[b], 这里需要维护一个单调递减序列

/*
//in: [3,4,5,7,2,9,0,1,2] 0-8 347<-014
	aS [-1,0,0,0,-1,0,-1,6,6]
	bS [0, 1,2,3, 4,5, 6,7,8]
	cS [4, 4,4,4, 6,6,-1,-1,-1]
	c' [1, 2,3,4,-1,6,-1,-1,-1]
	st [6,4,0
	//de [3,2,0]
	  [5,4,3,7,4
	  aS [-1,-1,-1,5,3]
	  descSeq [5,4,3
	  1,2,3,4,-1,6,-1,-1,-1
	  //in: [3,4,5,7,2,9,0,1,2]
	  //out: (0,1,4)
//out: */

class Solution{
	public:
		tuple<int,int,int> findLeastDictTuple(<vector>& nums){
			//利用st: 单调递减栈
			stack<int> st;
			//descSeq: 单调递减序列
			vector<int> descSeq;
			int n = nums.size();
			//最终的返回值
			tuple<int,int,int> res;
			vector<int> bcS(n,-1);//维护nums数组中对应的b和c的索引对
			
			//第一步利用单调栈找到b,c, 并将nums数组中存储其对应的每个最近的比其小的索引的数组(没找到的直接放-1)
			//第一步是找c
			//cS [4, 4,4,4, 6,6,-1,-1,-1]
			//c' [1, 2,3,4,-1,6,-1,-1,-1]
			for(int i = n - 1; i >= 0; i--){
				while(!st.empty() && nums[st.top()] >= nums[i]){
					st.pop();
				}
				if(!st.empty()){
					bcS[i] = st.top();
				}
				st.push(i);
			}
			
			//第二步利用单调递减序列结合二分查找来找到a
			//第二步是找a;找a和c的过程是独立进行的 二分查找进行的时候需要有序或者是单调性
			vector<int> aS;
			for(int j = 0; j < n; j++){
				if(descSeq.empty() || nums[descSeq[descSeq.size()-1]] >= nums[j]){
					aS.push_back(-1);
					descSeq.push_back(j);
				}
				else{
					int left = 0, right = descSeq.size()-1;
					while(left < right){
						int mid = left + (right - left)/2;
						if(nums[descSeq[mid]] < nums[j]){
							right = mid;
						}
						else if(nums[descSeq[mid]] >= nums[j]){
							left = mid + 1;
						}
					}
					aS.push_back(left);
				}
			}
			//aS和bcS就位
			vector<vector<int>> abc;
			for(int b = 0; b < n; b++){
				if(aS[b] == -1 || bcS[b] == -1){
					continue;
				}
				abc.push_back({aS[b],b,bcS[b]});
			}
			return *min_element(abc.begin(),abc.end());
		}
};