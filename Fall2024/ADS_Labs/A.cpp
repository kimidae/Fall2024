#include <iostream>
#include <deque>
#include <vector>

using namespace std;

vector<int> create_deque(int a) {
    deque<int> dq;
    vector<int> res_dq;
    
    for (int i = 1; i <= a; i++) {
        dq.push_front(i);
    }
    for (int i = 1; i <= a; i++) {
        if (!dq.empty()) {
            for (int j = 1; j <= i; j++) {
                dq.push_back(dq.front());
                dq.pop_front();
            }
                        if (dq.front() == i) {
                res_dq.push_back(i);
            }
        } else {
            if (res_dq.empty()) {
                res_dq.push_back(-1); 
            }
            while (!res_dq.empty()) {
                cout << res_dq.front() << " ";
                res_dq.erase(res_dq.begin()); 
            }
        }
    }
    
    return res_dq;
}

int main() {
    int n;
    cin >> n;
    vector<int> a(n);
    for (int i = 0; i < n; i++) {
        cin >> a[i];
    }

    for (int i = 0; i < n; i++) {
        vector<int> result = create_deque(a[i]);
        if (result.empty() || (result.size() == 1 && result[0] == -1)) {
            cout << "-1";
        } else {
            for (int num : result) {
                cout << num << " ";
            }
        }
        cout << endl;
    }
    
    return 0;
}
