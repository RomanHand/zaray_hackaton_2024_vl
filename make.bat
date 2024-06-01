pip install -r .\neiro\requirements.txt 
pyinstaller .\neiro\detector.py --onefile --noconsole --add-data=".\ultralytics\cfg\default.yaml:./ultralytics/cfg"
go build .\amnyam\cmd\amnyam\main.go -o C:/neiro/bin/back.exe
