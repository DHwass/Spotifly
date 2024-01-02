import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.song import SongSchema
from src.models.http_exceptions import *



songs_url = "http://localhost:8181/songs/"  # URL de l'API users (golang)

def get_song(id):
    response = requests.request(method="GET", url=songs_url+id)
    return response.json(), response.status_code

def get_songs():
    response = requests.request(method="GET", url=songs_url)
    return response.json(), response.status_code