pip install -r .\cutter\requirements.txt 
pip install -r .\neiro\requirements.txt 
pyinstaller .\cutter\cut.py --onefile --clean
copy .\cutter\dist\cut C:/neiro/bin/cut
pyinstaller .\neiro\detector.py --onefile --noconsole --add-data=".\ultralytics\cfg\default.yaml:./ultralytics/cfg"
copy .\cutter\dist\cut C:/neiro/bin/detector
go build amnyam/cmd/amnyam/main.go -o C:/neiro/bin/amnyam
