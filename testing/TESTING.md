
### How to run the testing:

1. Launch the server $ ./run-server.sh src/myconfig.ini (Port: 8088)
2. Go to '/testing' folder to run $ .testing.sh
3. The testing.sh sciprt will automatically run test cases and report the responses in 'response.output' file.


### For testing cases, :

1. BASIC FUNCTIONALITY FOR 200 CODE 
   
   - python test_header_body_set_200.py 
   - python test_file_kitten.py 
   - python test_file_ucsd.py

   - python test_dir.py 
   - python test_subdir.py 
   - python test_dir_without_slash.py

   - pythoh test_vali_extenstion_MIME.py 
   - python test_unvali_extension_MIME.py

2. CONCURRENCY and PIPELINING

   - ./pipeline.sh

   - python test_timeout.py  #Send first part of header, wait for 10sec, tnen send the rest part of header.

3. BASIC FUNCTIONALITY FOR NON-200 ERROR CODE

   - python test_missing_CRLF.py
   - python test_partical_CRLF.py
   - python test_header_no_space.py

   - python test_file_not_found_404.py
   - python test_urls_escape_docroot.py 
   - python test_out_of_docroot.py 

   - python test_empty_headerpath.py
   - python test_partical_header.py
   - python test_filepath_without_slash_beginning.py


#### CSE 124/224 Module 2 Project Milestone

    1. Name 1: Zhiqiang Sun 
    2. PID 1: A53304794 
    3. Github ID 1: BoscoSuen

    4. Name 2: Yifan Wang 
    5. PID 2: A53298382 
    6. Github ID 2: evawyf