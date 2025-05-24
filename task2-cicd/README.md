


Jenkins [새로운 Item] 추가
   - item name: beautifulstore
   - GitHub: https://github.com/leesunmi99/beautifulstore
   - sparse checkout: task2
   
```bash
cd task2
docker build -t beautifulstore:v1 .
docker rm -f beautifulstore || true
docker run -d --name beautifulstore -p 8080:8080 beautifulstore:v1
```


# 실습 과정



![image](https://github.com/user-attachments/assets/977edcf1-f455-4c98-b17c-8123a7f3bdaf)




![image](https://github.com/user-attachments/assets/dd9a2b49-f100-4141-becb-b766e9527b66)




<img width="1184" alt="image" src="https://github.com/user-attachments/assets/07010259-cad6-446f-9ad5-b726a2271088" />

<img width="1203" alt="image" src="https://github.com/user-attachments/assets/e32de41b-d875-4486-9c1b-16e232df46f8" />

# 결과 
![image](https://github.com/user-attachments/assets/d199072b-98be-4e47-acd7-d9a927237000)

![image](https://github.com/user-attachments/assets/cf3f5e03-55b4-447a-a3e1-8d8094c57152)
