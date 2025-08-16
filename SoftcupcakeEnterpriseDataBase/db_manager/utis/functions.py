import os
import json

def create_db(db_name, db_info):
    # db 디렉토리 생성
    os.makedirs(f'./{db_name}/tables', exist_ok=True)
    os.makedirs(f'./{db_name}/Archive', exist_ok=True)

    # db info파일 생성
    db_info_formatted = {
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
