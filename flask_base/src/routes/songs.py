import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.song import  SongUpdateSchema
from src.schemas.errors import *
import src.services.songs as songs_service

# from routes import songs  
songs = Blueprint(name="songs", import_name=__name__)

@songs.route('/<id>', methods=['GET'])
#@login_required
def get_song(id):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    print("id a get song",id)
    return songs_service.get_song(id)


@songs.route('/', methods=['GET'])
#@login_required
def get_songs():
    """
    ---
    get:
      description: Getting all songs
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Songs
            application/yaml:
              schema: Songs
        '401':

          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
    """
    return songs_service.get_songs()

@songs.route('/', methods=['POST'])
#@login_required
def post_song():
    """
    ---
    post:
      description: Creating a song
      requestBody:
        required: true
        content:
            application/json:
                schema: SongCreate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
    """
    print(request.json, "this is from route")
    return songs_service.create_song(request.json)

@songs.route('/<id>', methods=['PUT'])
#@login_required
def put_song(id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: SongUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
    """
    return songs_service.update_song(id, request.json)

@songs.route('/<id>', methods=['DELETE'])
#@login_required
def delete_song(id):
    """
    ---
    delete:
      description: Deleting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '204':
          description: No content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    return songs_service.delete_song(id)
#after having done everything for the songs, we will do the same for the ratings of the songs
@songs.route('/<id>/ratings', methods=['GET'])
#@login_required
def get_ratings(id):
    """
    ---
    get:
      description: Getting all ratings of a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Ratings
            application/yaml:
              schema: Ratings
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - ratings
    """
    return songs_service.get_ratings(id)
  
@songs.route('/<id>/ratings', methods=['POST'])
def add_rating(id):
        """
        ---
        post:
          description: Creating a rating
          parameters:
            - in: path
              name: id
              schema:
                type: uuidv4
              required: true
              description: UUID of song id
          requestBody:
            required: true
            content:
                application/json:
                    schema: RatingCreate
          responses:
            '201':
              description: Created
              content:
                application/json:
                  schema: Rating
                application/yaml:
                  schema: Rating
            '401':
              description: Unauthorized
              content:
                application/json:
                  schema: Unauthorized
                application/yaml:
                  schema: Unauthorized
            '422':
              description: Unprocessable entity
              content:
                application/json:
                  schema: UnprocessableEntity
                application/yaml:
                  schema: UnprocessableEntity
          tags:
              - ratings
        """
        return songs_service.create_rating(id, request.json)

@songs.route('/<id>/ratings/<id_rating>', methods=['GET'])
#@login_required
def get_rating(id,id_rating):
    """
    ---
    get:
      description: Getting a rating of a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
        - in: path
          name: id_rating
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - ratings
    """
    return songs_service.get_rating(id,id_rating)
@songs.route('/<id>/ratings/<id_rating>', methods=['PUT'])
#@login_required
def update_rating(id,id_rating):
    """
    ---
    put:
      description: Updating a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
        - in: path
          name: id_rating
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - ratings
    """
    return songs_service.update_rating(id,id_rating, request.json)
  

