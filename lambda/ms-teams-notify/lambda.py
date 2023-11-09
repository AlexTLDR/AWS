#!/usr/bin/python3.8
import urllib3
import json

http = urllib3.PoolManager()

def lambda_handler(event, context):
    # List of webhook URLs - replace with actual webhooks from MS Teams
    webhook_urls = ["https://your-webhook.webhook.office.com", 
                    "https://your-webhook.webhook.office.com"]

    for url in webhook_urls:
        msg = {
            "text": event['Records'][0]['Sns']['Message']
        }
        encoded_msg = json.dumps(msg).encode('utf-8')
        resp = http.request('POST', url, body=encoded_msg)
        print({
            "message": event['Records'][0]['Sns']['Message'],
            "status_code": resp.status,
            "response": resp.data
        })
