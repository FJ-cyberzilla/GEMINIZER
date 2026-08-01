import pytest
from unittest.mock import MagicMock
from src.geminizer_python.tasks.prompt_task import PromptTask
from src.geminizer_python.core.models import Context, TaskResult

def test_prompt_task_success() -> None:
    # Setup
    mock_client = MagicMock()
    mock_response = MagicMock()
    mock_response.text = "Enhanced cinematic prompt"
    mock_client.models.generate_content.return_value = mock_response
    
    task = PromptTask(client=mock_client)
    
    context = Context(data={"raw_prompt": "a sunset", "style": "cinematic"})
    
    # Execute
    result = task.execute(context)
    
    # Verify
    assert result.success is True
    assert result.data["enhanced_prompt"] == "Enhanced cinematic prompt"
    
    # Verify API call structure
    mock_client.models.generate_content.assert_called_once()
    args, kwargs = mock_client.models.generate_content.call_args
    assert "Style: cinematic" in kwargs["contents"]
    assert "User Prompt: a sunset" in kwargs["contents"]

def test_prompt_task_no_prompt() -> None:
    mock_client = MagicMock()
    task = PromptTask(client=mock_client)
    context = Context(data={})
    
    result = task.execute(context)
    
    assert result.success is False
    assert result.message == "No prompt provided."

def test_prompt_task_api_failure() -> None:
    # Setup
    mock_client = MagicMock()
    mock_client.models.generate_content.side_effect = Exception("API Error")
    
    task = PromptTask(client=mock_client)
    
    context = Context(data={"raw_prompt": "a sunset"})
    
    # Execute
    result = task.execute(context)
    
    # Verify
    assert result.success is False
    assert "Failed to enhance prompt" in result.message
