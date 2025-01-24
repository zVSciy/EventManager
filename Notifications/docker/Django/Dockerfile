FROM python:3.11

# Set the working directory
WORKDIR /app

COPY /frontend/requirements.txt .

# Copy the project files
COPY /frontend/frontend_notifications/. .

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# WORKDIR /app/frontend_notifications

# Expose the port the app runs on
EXPOSE 8090

# Command to run the application
CMD ["python", "manage.py", "runserver", "0.0.0.0:8090"]