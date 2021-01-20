
printf 'TESTING RESULT: \n\n' > response.output

printf '######## BASIC FUNCTIONALITY FOR 200 CODE ######## \n\n' >> response.output

python test_header_body_set_200.py 
python test_file_kitten.py
python test_file_ucsd.py

python test_dir.py 
python test_subdir.py 
python test_dir_without_slash.py 

pythoh test_vali_extenstion_MIME.py
python test_unvali_extension_MIME.py 

printf '######## CONCURRENCY Pipelining ######## \n\n' >> response.output

# python test_concurrency.py 
./pipeline.sh
# Send first part of header, wait for 10sec, tnen send the rest part of header. 
python test_timeout.py 
python test_connClose.py

printf '######## BASIC FUNCTIONALITY FOR NON-200 ERROR CODE ######## \n\n ' >> response.output

python test_missing_CRLF.py
python test_partical_CRLF.py
python test_header_no_space.py

python test_file_not_found_404.py
python test_urls_escape_docroot.py 
python test_out_of_docroot.py 

python test_filepath_without_slash_beginning.py


