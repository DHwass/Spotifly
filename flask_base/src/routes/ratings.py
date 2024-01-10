import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.rating import  RatingUpdateSchema
from src.schemas.errors import *
import src.services.ratings as ratings_service

# from routes import ratings  
ratings = Blueprint(name="ratings", import_name=__name__)
@ratings.route('/<id>', methods=['POST'])
#@login_required
def post_rating(id):
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
            schema: RatingUpdate
          application/yaml:
            schema: RatingUpdate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '400':
          description: Bad request
          content:
            application/json:
              schema: BadRequest
            application/yaml:
              schema: BadRequest
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

  
    return ratings_service.create_rating(id,request.json)

@ratings.route('/<id>', methods=['GET'])
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
    return ratings_service.get_ratings(id)