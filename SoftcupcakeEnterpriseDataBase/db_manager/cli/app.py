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
        "Archive Storage Directory": db_info["Archive Storage Directory"],
        "server port": db_info["server port"]
    }

    print(db_info_formatted)

    # json 작성
    with open(f"./{db_name}/info.json", "w", encoding="utf-8") as f:
        json.dump(db_info_formatted, f, indent=4, ensure_ascii=False)


# db정보
db_info = {
    "version": 0.01,
    "Is beta": "Beta",   # 여기 키값은 그대로 두되, 저장 시 "Is beta"로 바꿔서 씀
    "Archive Storage Directory": "./Archive",
    "server port": 8000
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