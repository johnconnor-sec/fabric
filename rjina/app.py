import requests
from urllib.parse import urlencode
import argparse

def make_request(input_str: str) -> requests.Response:
  """Makes a GET request to the Jina API with URL encoded input.

  Args:
    input_str: The input string to be URL encoded and appended to the URL.

  Returns:
    The response object from the Jina API.
  """
  url = f"https://r.jina.ai/{input_str}"
  print(url)
  response = requests.get(url)
  return response

if __name__ == "__main__":
  parser = argparse.ArgumentParser(description="Make a request to the Jina API")
  parser.add_argument("input", type=str, help="The url of the page to read.")
  args = parser.parse_args()

  response = make_request(args.input)
  print(response.text)
