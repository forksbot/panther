# Panther is a Cloud-Native SIEM for the Modern Security Team.
# Copyright (C) 2020 Panther Labs Inc
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

from unittest import TestCase, mock

from ..src import EventMatch
from ..src.engine import Engine


class TestEngine(TestCase):

    def test_loading_rules(self) -> None:
        analysis_api = mock.MagicMock()
        analysis_api.get_enabled_rules.return_value = [
            {
                'id': 'rule_id',
                'resourceTypes': ['log'],
                'body': 'def rule(event):\n\treturn True',
                'versionId': 'version'
            }
        ]
        engine = Engine(analysis_api)
        self.assertEqual(len(engine.log_type_to_rules), 1)
        self.assertEqual(len(engine.log_type_to_rules['log']), 1)
        self.assertEqual(engine.log_type_to_rules['log'][0].rule_id, 'rule_id')

    def test_analyse_many_rules(self) -> None:
        analysis_api = mock.MagicMock()
        analysis_api.get_enabled_rules.return_value = [
            {
                'id': 'rule_id_1',
                'resourceTypes': ['log'],
                'body': 'def rule(event):\n\treturn True',
                'versionId': 'version',
                'dedupPeriodMinutes': 120
            },  # This rule should match the event
            {
                'id': 'rule_id_2',
                'resourceTypes': ['log'],
                'body': 'def rule(event):\n\treturn False',
                'versionId': 'version'
            }  # This rule shouldn't match the event
        ]
        engine = Engine(analysis_api)
        result = engine.analyze('log', {})

        expected_event_matches = [
            EventMatch(
                rule_id='rule_id_1',
                rule_version='version',
                log_type='log',
                dedup='defaultDedupString:rule_id_1',
                dedup_period_mins=120,
                event={}
            )
        ]
        self.assertEqual(result, expected_event_matches)

    def test_analyse_many_rules_one_throws_exception(self) -> None:
        analysis_api = mock.MagicMock()
        analysis_api.get_enabled_rules.return_value = [
            {
                'id': 'rule_id_1',
                'resourceTypes': ['log'],
                'body': 'def rule(event):\n\treturn True',
                'versionId': 'version'
            }, {
                'id': 'rule_id_2',
                'resourceTypes': ['log'],
                'body': 'def rule(event):\n\traise Exception()',
                'versionId': 'version'
            }, {
                'id': 'rule_id_3',
                'resourceTypes': ['log'],
                'body': 'def rule(event):\n\treturn True',
                'versionId': 'version'
            }
        ]
        engine = Engine(analysis_api)
        result = engine.analyze('log', {})

        expected_event_matches = [
            EventMatch(
                rule_id='rule_id_1',
                rule_version='version',
                log_type='log',
                dedup='defaultDedupString:rule_id_1',
                event={},
                dedup_period_mins=60
            ),
            EventMatch(
                rule_id='rule_id_3',
                rule_version='version',
                log_type='log',
                dedup='defaultDedupString:rule_id_3',
                event={},
                dedup_period_mins=60
            )
        ]

        self.assertEqual(result, expected_event_matches)
