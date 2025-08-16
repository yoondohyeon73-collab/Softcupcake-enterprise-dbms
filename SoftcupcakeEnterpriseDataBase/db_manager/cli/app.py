import os
import json

def create_db(db_name, db_info):
    # db 디렉토리 생성
    os.makedirs(f'./{db_name}/tables', exist_ok=True)
    os.makedirs(f'./{db_name}/Archive', exist_ok=True)

    # db info파일 생성
    db_info_formatted = {
        "db name" : db_name,
        "dbms version": {
            "version": db_info["version"],
            "Is beta": db_info["Is beta"]   # 오타 수정: Is bata → Is beta
        },
        "server port": db_info["server port"]
    }

    print(db_info_formatted)

    # json 작성
    with open(f"./{db_name}/info.json", "w", encoding="utf-8") as f:
        json.dump(db_info_formatted, f, indent=4, ensure_ascii=False)

#db버전
db_version = {
    "version" : 0.01,
    "is Beta" : "Beta"
}

#메뉴얼 출력
print_arg = '''

Softcupcake Enterprise Database
Copyright © 2025 윤도현. All rights reserved.

--------------------------------------------

Command Manual

When creating a table 
-> Enter 'create' and follow the command
If you want to run the DB
-> Enter 'serve' and follow the command

'''

print(print_arg)

while True:
    user_cmd = input("CMD : ")

    #db생성
    if user_cmd == 'create' or user_cmd == 'Create':

        #db정보 만들기
        db_info = {}

        #db이름/실행할 포트 입력받기
        db_name = input("Enter db name : ")
        server_port = input("Enter the port to run the server : ")

        #info정의
        db_info["name"] = db_name
        db_info["server port"] = server_port
        db_info["version"] = db_version["version"]
        db_info["Is beta"] = db_version["is Beta"]

        #db정보 출력
        print('Check db information\n')
        print(db_info)
        print('\nCreate db (y/n)\n')
        
        #db생성 (y/n)에 따라 실행
        yorn = input('')
        if yorn == 'y' or yorn == 'Y':
            create_db(db_info["name"], db_info)
        else:
            print('End db creation\n')
            db_info = {}
            
