# python image
FROM python:3.12

# python dependencies
WORKDIR /usr/local/app
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

# python script
COPY src ./src

# run python script, need to pass a value as argument during run time
ENTRYPOINT ["python", "src/cos-value.py"]