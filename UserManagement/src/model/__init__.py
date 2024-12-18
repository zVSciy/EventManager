from . import models, database, security

models.Base.metadata.create_all(bind=database.engine)