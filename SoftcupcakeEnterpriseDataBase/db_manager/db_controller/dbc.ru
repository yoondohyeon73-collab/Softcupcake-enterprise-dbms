#일단 이 db컨트롤러는 modules속 db컨트롤러와 다른 컨트롤러로 
#CLI용 컨트롤러임
#보통 파일입출력이 주된 기능임

require 'fileutils'

#db정보 구조체
db_info = Struct.new(:db_name, :autoArchiving)

#db생성
def create_db(db_info) 

  #db디렉토리 생성
  FileUtils.mkdir_p("./ROOT/"+db_info.db_name)

  #정보파일 json내용 버퍼
  buffer = '{"DB_NAME" : "'+db_info.db_name+'", "AUTO_ARCH" : "'db_info.autoArchiving+'"}'

  #정보 파일 생성
  File.open("./ROOT/"+db_info.db_name+"/info.json", "w") do |f|
  f.puts buffer
  end

end 