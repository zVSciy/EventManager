from . import models, database, config, security

models.Base.metadata.create_all(bind=database.engine)