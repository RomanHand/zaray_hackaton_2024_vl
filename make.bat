pip install -r .\cutter\requirements.txt 
pip install -r .\neiro\requirements.txt 
pyinstaller .\cutter\cut.py --onefile --clean
copy .\cutter\dist\cut C:/neiro/bin/cut
pyinstaller .\neiro\detector.py --onefile --clean
copy .\cutter\dist\cut C:/neiro/bin/detector

