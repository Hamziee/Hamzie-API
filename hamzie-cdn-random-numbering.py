import os

directory = 'C:/Users/Hamza/Pictures/xiaojie/'
extension = '.jpg'

# Get a list of all files in the directory
files = os.listdir(directory)

# Sort the files alphabetically
files.sort()

# Initialize a counter for numbering the files
counter = 1

# Iterate over the files
for file in files:
    # Check if the file is a JPG file
    if file.endswith(extension):
        # Generate the new file name with the counter
        new_name = str(counter) + extension
        
        # Check if the new file name already exists
        while os.path.exists(os.path.join(directory, new_name)):
            counter += 1
            new_name = str(counter) + extension
        
        # Rename the file
        os.rename(os.path.join(directory, file), os.path.join(directory, new_name))
        
        # Increment the counter
        counter += 1